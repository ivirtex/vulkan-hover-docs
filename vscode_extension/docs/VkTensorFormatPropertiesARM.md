# VkTensorFormatPropertiesARM(3) Manual Page

## Name

VkTensorFormatPropertiesARM - Structure specifying properties of a format used to describe tensor elements



## [](#_c_specification)C Specification

The [VkTensorFormatPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorFormatPropertiesARM.html) structure describes properties of a [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) when that format is used to describe tensor elements. These properties, like those of [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatProperties2.html), are independent of any particular tensor.

The [VkTensorFormatPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorFormatPropertiesARM.html) structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkTensorFormatPropertiesARM {
    VkStructureType          sType;
    const void*              pNext;
    VkFormatFeatureFlags2    optimalTilingTensorFeatures;
    VkFormatFeatureFlags2    linearTilingTensorFeatures;
} VkTensorFormatPropertiesARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `linearTilingTensorFeatures` is a bitmask of [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html) specifying features supported by tensors created with a `tiling` parameter of `VK_TENSOR_TILING_LINEAR_ARM`.
- `optimalTilingTensorFeatures` is a bitmask of [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlagBits2.html) specifying features supported by tensors created with a `tiling` parameter of `VK_TENSOR_TILING_OPTIMAL_ARM`.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkTensorFormatPropertiesARM-sType-sType)VUID-VkTensorFormatPropertiesARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_TENSOR_FORMAT_PROPERTIES_ARM`

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTensorFormatPropertiesARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0