# VkExternalMemoryFeatureFlagBitsNV(3) Manual Page

## Name

VkExternalMemoryFeatureFlagBitsNV - Bitmask specifying external memory
features



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html)::`externalMemoryFeatures`,
indicating properties of the external memory handle type, are:

``` c
// Provided by VK_NV_external_memory_capabilities
typedef enum VkExternalMemoryFeatureFlagBitsNV {
    VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV = 0x00000001,
    VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV = 0x00000002,
    VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV = 0x00000004,
} VkExternalMemoryFeatureFlagBitsNV;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV` specifies that
  external memory of the specified type **must** be created as a
  dedicated allocation when used in the manner specified.

- `VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV` specifies that the
  implementation supports exporting handles of the specified type.

- `VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV` specifies that the
  implementation supports importing handles of the specified type.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_capabilities.html),
[VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html),
[VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryFeatureFlagsNV.html),
[vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExternalMemoryFeatureFlagBitsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
