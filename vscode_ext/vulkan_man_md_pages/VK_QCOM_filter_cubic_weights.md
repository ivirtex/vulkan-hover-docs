# VK_QCOM_filter_cubic_weights(3) Manual Page

## Name

VK_QCOM_filter_cubic_weights - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

520

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_EXT_filter_cubic](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_filter_cubic.html)  

## <a href="#_contact" class="anchor"></a>Contact

- Matthew Netsch <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_filter_cubic_weights%5D%20@mnetsch%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_filter_cubic_weights%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>mnetsch</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-06-23

**Contributors**  
- Jeff Leger, Qualcomm Technologies, Inc.

- Jonathan Wicks, Qualcomm Technologies, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension extends cubic filtering by adding the ability to select a
set of weights. Without this extension, the weights used in cubic
filtering are limited to those corresponding to a Catmull-Rom spline.
This extension adds support for 3 additional spline weights.

This extension adds a new structure that **can** be added to the `pNext`
chain of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html) that **can** be
used to specify which set of cubic weights are used in cubic filtering.
A similar structure can be added to the `pNext` chain of
[VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html) to specify cubic weights used
in a blit operation.

With this extension weights corresponding to the following additional
splines can be selected for cubic filtered sampling and blits:

- Zero Tangent Cardinal

- B-Spline

- Mitchell-Netravali

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html):

  - [VkBlitImageCubicWeightsInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageCubicWeightsInfoQCOM.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCubicWeightsFeaturesQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCubicWeightsFeaturesQCOM.html)

- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html):

  - [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCubicFilterWeightsQCOM.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_QCOM_FILTER_CUBIC_WEIGHTS_EXTENSION_NAME`

- `VK_QCOM_FILTER_CUBIC_WEIGHTS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_BLIT_IMAGE_CUBIC_WEIGHTS_INFO_QCOM`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUBIC_WEIGHTS_FEATURES_QCOM`

  - `VK_STRUCTURE_TYPE_SAMPLER_CUBIC_WEIGHTS_CREATE_INFO_QCOM`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2023-06-23 (jleger)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_QCOM_filter_cubic_weights"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
