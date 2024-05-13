# VkExternalImageFormatPropertiesNV(3) Manual Page

## Name

VkExternalImageFormatPropertiesNV - Structure specifying external image
format properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkExternalImageFormatPropertiesNV` structure is defined as:

``` c
// Provided by VK_NV_external_memory_capabilities
typedef struct VkExternalImageFormatPropertiesNV {
    VkImageFormatProperties              imageFormatProperties;
    VkExternalMemoryFeatureFlagsNV       externalMemoryFeatures;
    VkExternalMemoryHandleTypeFlagsNV    exportFromImportedHandleTypes;
    VkExternalMemoryHandleTypeFlagsNV    compatibleHandleTypes;
} VkExternalImageFormatPropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `imageFormatProperties` will be filled in as when calling
  [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html),
  but the values returned **may** vary depending on the external handle
  type requested.

- `externalMemoryFeatures` is a bitmask of
  [VkExternalMemoryFeatureFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagBitsNV.html),
  indicating properties of the external memory handle type
  ([vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType`)
  being queried, or 0 if the external memory handle type is 0.

- `exportFromImportedHandleTypes` is a bitmask of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  containing a bit set for every external handle type that **may** be
  used to create memory from which the handles of the type specified in
  [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType`
  **can** be exported, or 0 if the external memory handle type is 0.

- `compatibleHandleTypes` is a bitmask of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  containing a bit set for every external handle type that **may** be
  specified simultaneously with the handle type specified by
  [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType`
  when calling [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html), or 0 if the
  external memory handle type is 0. `compatibleHandleTypes` will always
  contain
  [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)::`externalHandleType`

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_capabilities.html),
[VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagsNV.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html),
[VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html),
[vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalImageFormatPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
