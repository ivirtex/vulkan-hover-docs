# VkPhysicalDevice16BitStorageFeatures(3) Manual Page

## Name

VkPhysicalDevice16BitStorageFeatures - Structure describing features supported by VK\_KHR\_16bit\_storage



## [](#_c_specification)C Specification

The [VkPhysicalDevice16BitStorageFeatures](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice16BitStorageFeatures.html) structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkPhysicalDevice16BitStorageFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           storageBuffer16BitAccess;
    VkBool32           uniformAndStorageBuffer16BitAccess;
    VkBool32           storagePushConstant16;
    VkBool32           storageInputOutput16;
} VkPhysicalDevice16BitStorageFeatures;
```

or the equivalent

```c++
// Provided by VK_KHR_16bit_storage
typedef VkPhysicalDevice16BitStorageFeatures VkPhysicalDevice16BitStorageFeaturesKHR;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.

## [](#_description)Description

- []()`storageBuffer16BitAccess` specifies whether objects in the `StorageBuffer`, `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer` storage class with the `Block` decoration **can** have 16-bit integer and 16-bit floating-point members. If this feature is not enabled, 16-bit integer or 16-bit floating-point members **must** not be used in such objects unless [`storageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extension-features-storageBuffer8BitAccess) or [`uniformAndStorageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extension-features-uniformAndStorageBuffer8BitAccess) are enabled or they are accessed in 32-bit multiples if [`shaderUntypedPointers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderUntypedPointers) is enabled. This also specifies whether shader modules **can** declare the `StorageBuffer16BitAccess` capability.
- []()`uniformAndStorageBuffer16BitAccess` specifies whether objects in the `Uniform` storage class with the `Block` decoration **can** have 16-bit integer and 16-bit floating-point members. If this feature is not enabled, 16-bit integer or 16-bit floating-point members **must** not be used in such objects unless [`uniformAndStorageBuffer8BitAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extension-features-uniformAndStorageBuffer8BitAccess) are enabled or they are accessed in 32-bit multiples if [`shaderUntypedPointers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderUntypedPointers) is enabled. This also specifies whether shader modules **can** declare the `UniformAndStorageBuffer16BitAccess` capability.
- []()`storagePushConstant16` specifies whether objects in the `PushConstant` storage class **can** have 16-bit integer and 16-bit floating-point members. If this feature is not enabled, 16-bit integer or floating-point members **must** not be used in such objects unless [`storagePushConstant8`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extension-features-storagePushConstant8) are enabled or they are accessed in 32-bit multiples if [`shaderUntypedPointers`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderUntypedPointers) is enabled. This also specifies whether shader modules **can** declare the `StoragePushConstant16` capability.
- []()`storageInputOutput16` specifies whether objects in the `Input` and `Output` storage classes **can** have 16-bit integer and 16-bit floating-point members. If this feature is not enabled, 16-bit integer or 16-bit floating-point members **must** not be used in such objects. This also specifies whether shader modules **can** declare the `StorageInputOutput16` capability.

If the `VkPhysicalDevice16BitStorageFeatures` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDevice16BitStorageFeatures`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDevice16BitStorageFeatures-sType-sType)VUID-VkPhysicalDevice16BitStorageFeatures-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDevice16BitStorageFeatures).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0