# VkBlitImageCubicWeightsInfoQCOM(3) Manual Page

## Name

VkBlitImageCubicWeightsInfoQCOM - Structure specifying image blit cubic weight info



## [](#_c_specification)C Specification

The `VkBlitImageCubicWeightsInfoQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_filter_cubic_weights
typedef struct VkBlitImageCubicWeightsInfoQCOM {
    VkStructureType             sType;
    const void*                 pNext;
    VkCubicFilterWeightsQCOM    cubicWeights;
} VkBlitImageCubicWeightsInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `cubicWeights` is a [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html) value controlling cubic filter weights for the blit.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkBlitImageCubicWeightsInfoQCOM-sType-sType)VUID-VkBlitImageCubicWeightsInfoQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BLIT_IMAGE_CUBIC_WEIGHTS_INFO_QCOM`
- [](#VUID-VkBlitImageCubicWeightsInfoQCOM-cubicWeights-parameter)VUID-VkBlitImageCubicWeightsInfoQCOM-cubicWeights-parameter  
  `cubicWeights` **must** be a valid [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html) value

## [](#_see_also)See Also

[VK\_KHR\_copy\_commands2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_commands2.html), [VK\_QCOM\_filter\_cubic\_weights](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_filter_cubic_weights.html), [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBlitImageCubicWeightsInfoQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0