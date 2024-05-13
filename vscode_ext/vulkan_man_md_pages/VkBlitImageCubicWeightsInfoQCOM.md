# VkBlitImageCubicWeightsInfoQCOM(3) Manual Page

## Name

VkBlitImageCubicWeightsInfoQCOM - Structure specifying image blit cubic
weight info



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBlitImageCubicWeightsInfoQCOM` structure is defined as:

``` c
// Provided by VK_QCOM_filter_cubic_weights
typedef struct VkBlitImageCubicWeightsInfoQCOM {
    VkStructureType             sType;
    const void*                 pNext;
    VkCubicFilterWeightsQCOM    cubicWeights;
} VkBlitImageCubicWeightsInfoQCOM;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `cubicWeights` is a
  [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html) value
  controlling cubic filter weights for the blit.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkBlitImageCubicWeightsInfoQCOM-sType-sType"
  id="VUID-VkBlitImageCubicWeightsInfoQCOM-sType-sType"></a>
  VUID-VkBlitImageCubicWeightsInfoQCOM-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BLIT_IMAGE_CUBIC_WEIGHTS_INFO_QCOM`

- <a href="#VUID-VkBlitImageCubicWeightsInfoQCOM-cubicWeights-parameter"
  id="VUID-VkBlitImageCubicWeightsInfoQCOM-cubicWeights-parameter"></a>
  VUID-VkBlitImageCubicWeightsInfoQCOM-cubicWeights-parameter  
  `cubicWeights` **must** be a valid
  [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_QCOM_filter_cubic_weights](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QCOM_filter_cubic_weights.html),
[VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBlitImageCubicWeightsInfoQCOM"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
