# VkPhysicalDevice8BitStorageFeatures(3) Manual Page

## Name

VkPhysicalDevice8BitStorageFeatures - Structure describing features
supported by VK_KHR_8bit_storage



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDevice8BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice8BitStorageFeatures.html)
structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDevice8BitStorageFeatures {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           storageBuffer8BitAccess;
    VkBool32           uniformAndStorageBuffer8BitAccess;
    VkBool32           storagePushConstant8;
} VkPhysicalDevice8BitStorageFeatures;
```

or the equivalent

``` c
// Provided by VK_KHR_8bit_storage
typedef VkPhysicalDevice8BitStorageFeatures VkPhysicalDevice8BitStorageFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-storageBuffer8BitAccess"></span>
  `storageBuffer8BitAccess` indicates whether objects in the
  `StorageBuffer`, `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer`
  storage class with the `Block` decoration **can** have 8-bit integer
  members. If this feature is not enabled, 8-bit integer members
  **must** not be used in such objects. This also indicates whether
  shader modules **can** declare the `StorageBuffer8BitAccess`
  capability.

- <span id="extension-features-uniformAndStorageBuffer8BitAccess"></span>
  `uniformAndStorageBuffer8BitAccess` indicates whether objects in the
  `Uniform` storage class with the `Block` decoration **can** have 8-bit
  integer members. If this feature is not enabled, 8-bit integer members
  **must** not be used in such objects. This also indicates whether
  shader modules **can** declare the `UniformAndStorageBuffer8BitAccess`
  capability.

- <span id="extension-features-storagePushConstant8"></span>
  `storagePushConstant8` indicates whether objects in the `PushConstant`
  storage class **can** have 8-bit integer members. If this feature is
  not enabled, 8-bit integer members **must** not be used in such
  objects. This also indicates whether shader modules **can** declare
  the `StoragePushConstant8` capability.

If the `VkPhysicalDevice8BitStorageFeatures` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevice8BitStorageFeatures` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDevice8BitStorageFeatures-sType-sType"
  id="VUID-VkPhysicalDevice8BitStorageFeatures-sType-sType"></a>
  VUID-VkPhysicalDevice8BitStorageFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_8BIT_STORAGE_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_8bit_storage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_8bit_storage.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevice8BitStorageFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
