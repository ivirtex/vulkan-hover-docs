# VkPhysicalDevice16BitStorageFeatures(3) Manual Page

## Name

VkPhysicalDevice16BitStorageFeatures - Structure describing features
supported by VK_KHR_16bit_storage



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkPhysicalDevice16BitStorageFeatures](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice16BitStorageFeatures.html)
structure is defined as:

``` c
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

``` c
// Provided by VK_KHR_16bit_storage
typedef VkPhysicalDevice16BitStorageFeatures VkPhysicalDevice16BitStorageFeaturesKHR;
```

## <a href="#_members" class="anchor"></a>Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-features-storageBuffer16BitAccess"></span>
  `storageBuffer16BitAccess` specifies whether objects in the
  `StorageBuffer`, `ShaderRecordBufferKHR`, or `PhysicalStorageBuffer`
  storage class with the `Block` decoration **can** have 16-bit integer
  and 16-bit floating-point members. If this feature is not enabled,
  16-bit integer or 16-bit floating-point members **must** not be used
  in such objects. This also specifies whether shader modules **can**
  declare the `StorageBuffer16BitAccess` capability.

- <span id="extension-features-uniformAndStorageBuffer16BitAccess"></span>
  `uniformAndStorageBuffer16BitAccess` specifies whether objects in the
  `Uniform` storage class with the `Block` decoration **can** have
  16-bit integer and 16-bit floating-point members. If this feature is
  not enabled, 16-bit integer or 16-bit floating-point members **must**
  not be used in such objects. This also specifies whether shader
  modules **can** declare the `UniformAndStorageBuffer16BitAccess`
  capability.

- <span id="extension-features-storagePushConstant16"></span>
  `storagePushConstant16` specifies whether objects in the
  `PushConstant` storage class **can** have 16-bit integer and 16-bit
  floating-point members. If this feature is not enabled, 16-bit integer
  or floating-point members **must** not be used in such objects. This
  also specifies whether shader modules **can** declare the
  `StoragePushConstant16` capability.

- <span id="extension-features-storageInputOutput16"></span>
  `storageInputOutput16` specifies whether objects in the `Input` and
  `Output` storage classes **can** have 16-bit integer and 16-bit
  floating-point members. If this feature is not enabled, 16-bit integer
  or 16-bit floating-point members **must** not be used in such objects.
  This also specifies whether shader modules **can** declare the
  `StorageInputOutput16` capability.

If the `VkPhysicalDevice16BitStorageFeatures` structure is included in
the `pNext` chain of the
[VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html) structure
passed to
[vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFeatures2.html), it is
filled in to indicate whether each corresponding feature is supported.
`VkPhysicalDevice16BitStorageFeatures` **can** also be used in the
`pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html) to
selectively enable these features.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDevice16BitStorageFeatures-sType-sType"
  id="VUID-VkPhysicalDevice16BitStorageFeatures-sType-sType"></a>
  VUID-VkPhysicalDevice16BitStorageFeatures-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_16BIT_STORAGE_FEATURES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDevice16BitStorageFeatures"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
