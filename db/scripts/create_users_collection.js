// use userDB

db.createCollection('users', {
    validator: {
        $and: [
            {
                $expr: {
                    $lt: ['$birthDate', '$joinedDate'],
                },
            },
            {
                $jsonSchema: {
                    bsonType: 'object',
                    required: [
                        '_id',
                        'userName',
                        'email',
                        'password',
                        'birthDate',
                    ],
                    properties: {
                        _id: {
                            bsonType: 'objectId',
                        },
                        userName: {
                            bsonType: 'string',
                            description:
                                'Username is required and should be a string',
                            minLength: 3,
                            maxLength: 15,
                        },
                        email: {
                            bsonType: 'string',
                            pattern:
                                '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$',
                            description:
                                'Email is required and should be a string',
                        },
                        password: {
                            bsonType: 'string',
                            description:
                                'Password is required and should be a string',
                        },
                        profilePicture: {
                            bsonType: 'string',
                            description:
                                'Profile picture should be an URL string',
                        },
                        bio: {
                            bsonType: 'string',
                            maxLength: 100,
                            description: 'Bio should be a string',
                        },
                        birthDate: {
                            bsonType: 'date',
                            description:
                                'Birth date is required and should be a date',
                        },
                        joinedDate: {
                            bsonType: 'date',
                            description:
                                'Joined date is required and should be a date',
                        },
                        activated: {
                            bsonType: 'bool',
                            description: 'Activated should be a boolean',
                        },
                        followers: {
                            bsonType: 'array',
                            items: {
                                bsonType: 'objectId',
                            },
                            description:
                                'Followers should be an array of objectIds',
                        },
                        following: {
                            bsonType: 'array',
                            items: {
                                bsonType: 'objectId',
                            },
                            description:
                                'Following should be an array of objectIds',
                        },
                    },
                    additionalProperties: false,
                },
            },
        ],
    },
});

db.users.createIndex({ email: 1 }, { unique: true });

// db.runCommand({
//     collMod: 'users',
//     validator: {
//         $and: [
//             {
//                 $expr: {
//                     $lt: ['$birthDate', '$joinedDate'],
//                 },
//             },
//             {
//                 $jsonSchema: {
//                     bsonType: 'object',
//                     required: [
//                         '_id',
//                         'userName',
//                         'email',
//                         'password',
//                         'birthDate',
//                     ],
//                     properties: {
//                         _id: {
//                             bsonType: 'objectId',
//                         },
//                         userName: {
//                             bsonType: 'string',
//                             description:
//                                 'Username is required and should be a string',
//                             minLength: 3,
//                             maxLength: 15,
//                         },
//                         email: {
//                             bsonType: 'string',
//                             pattern:
//                                 '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$',
//                             description:
//                                 'Email is required and should be a string',
//                         },
//                         password: {
//                             bsonType: 'string',
//                             description:
//                                 'Password is required and should be a string',
//                         },
//                         profilePicture: {
//                             bsonType: 'string',
//                             description:
//                                 'Profile picture should be an URL string',
//                         },
//                         bio: {
//                             bsonType: 'string',
//                             maxLength: 100,
//                             description: 'Bio should be a string',
//                         },
//                         birthDate: {
//                             bsonType: 'date',
//                             description:
//                                 'Birth date is required and should be a date',
//                         },
//                         joinedDate: {
//                             bsonType: 'date',
//                             description:
//                                 'Joined date is required and should be a date',
//                         },
//                         followers: {
//                             bsonType: 'array',
//                             items: {
//                                 bsonType: 'objectId',
//                             },
//                             description:
//                                 'Followers should be an array of objectIds',
//                         },
//                         following: {
//                             bsonType: 'array',
//                             items: {
//                                 bsonType: 'objectId',
//                             },
//                             description:
//                                 'Following should be an array of objectIds',
//                         },
//                     },
//                     additionalProperties: false,
//                 },
//             },
//         ],
//     },
// });
