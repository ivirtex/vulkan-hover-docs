# VkExternalImageFormatPropertiesNV(3) Manual Page

## Name

VkExternalImageFormatPropertiesNV - Structure specifying external image format properties



## [](#_c_specification)C Specification

The `VkExternalImageFormatPropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_external_memory_capabilities
typedef struct VkExternalImageFormatPropertiesNV {
    VkImageFormatProperties              imageFormatProperties;
    VkExternalMemoryFeatureFlagsNV       externalMemoryFeatures;
    VkExternalMemoryHandleTypeFlagsNV    exportFromImportedHandleTypes;
    VkExternalMemoryHandleTypeFlagsNV    compatibleHandleTypes;
} VkExternalImageFormatPropertiesNV;
```

## [](#_members)Members

- `imageFormatProperties` will be filled in as when calling [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties.html), but the values returned **may** vary depending on the external handle type requested.
- `externalMemoryFeatures` is a bitmask of [VkExternalMemoryFeatureFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagBitsNV.html), indicating properties of the external memory handle type ([vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType`) being queried, or 0 if the external memory handle type is 0.
- `exportFromImportedHandleTypes` is a bitmask of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) containing a bit set for every external handle type that **may** be used to create memory from which the handles of the type specified in [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType` **can** be exported, or 0 if the external memory handle type is 0.
- `compatibleHandleTypes` is a bitmask of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) containing a bit set for every external handle type that **may** be specified simultaneously with the handle type specified by [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType` when calling [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAllocateMemory.html), or 0 if the external memory handle type is 0. `compatibleHandleTypes` will always contain [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType`

## [](#_description)Description

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_capabilities.html), [VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagsNV.html), [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsNV.html), [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatProperties.html), [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalImageFormatPropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0