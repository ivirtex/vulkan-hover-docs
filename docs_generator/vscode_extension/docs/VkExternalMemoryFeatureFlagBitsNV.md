# VkExternalMemoryFeatureFlagBitsNV(3) Manual Page

## Name

VkExternalMemoryFeatureFlagBitsNV - Bitmask specifying external memory features



## [](#_c_specification)C Specification

Bits which **can** be set in [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html)::`externalMemoryFeatures`, indicating properties of the external memory handle type, are:

```c++
// Provided by VK_NV_external_memory_capabilities
typedef enum VkExternalMemoryFeatureFlagBitsNV {
    VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV = 0x00000001,
    VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV = 0x00000002,
    VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV = 0x00000004,
} VkExternalMemoryFeatureFlagBitsNV;
```

## [](#_description)Description

- `VK_EXTERNAL_MEMORY_FEATURE_DEDICATED_ONLY_BIT_NV` specifies that external memory of the specified type **must** be created as a dedicated allocation when used in the manner specified.
- `VK_EXTERNAL_MEMORY_FEATURE_EXPORTABLE_BIT_NV` specifies that the implementation supports exporting handles of the specified type.
- `VK_EXTERNAL_MEMORY_FEATURE_IMPORTABLE_BIT_NV` specifies that the implementation supports importing handles of the specified type.

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_capabilities.html), [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html), [VkExternalMemoryFeatureFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryFeatureFlagsNV.html), [vkGetPhysicalDeviceExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceExternalImageFormatPropertiesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalMemoryFeatureFlagBitsNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0