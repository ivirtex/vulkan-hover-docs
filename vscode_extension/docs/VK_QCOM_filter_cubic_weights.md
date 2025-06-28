# VK\_QCOM\_filter\_cubic\_weights(3) Manual Page

## Name

VK\_QCOM\_filter\_cubic\_weights - device extension



## [](#_registered_extension_number)Registered Extension Number

520

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_filter_cubic_weights%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_filter_cubic_weights%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-06-23

**Contributors**

- Jeff Leger, Qualcomm Technologies, Inc.
- Jonathan Wicks, Qualcomm Technologies, Inc.

## [](#_description)Description

This extension extends cubic filtering by adding the ability to select a set of weights. Without this extension, the weights used in cubic filtering are limited to those corresponding to a Catmull-Rom spline. This extension adds support for 3 additional spline weights.

This extension adds a new structure that **can** be added to the `pNext` chain of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html) that **can** be used to specify which set of cubic weights are used in cubic filtering. A similar structure can be added to the `pNext` chain of [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2.html) to specify cubic weights used in a blit operation.

With this extension weights corresponding to the following additional splines can be selected for cubic filtered sampling and blits:

- Zero Tangent Cardinal
- B-Spline
- Mitchell-Netravali

## [](#_new_structures)New Structures

- Extending [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageInfo2.html):
  
  - [VkBlitImageCubicWeightsInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBlitImageCubicWeightsInfoQCOM.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCubicWeightsFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCubicWeightsFeaturesQCOM.html)
- Extending [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html):
  
  - [VkSamplerCubicWeightsCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCubicWeightsCreateInfoQCOM.html)

## [](#_new_enums)New Enums

- [VkCubicFilterWeightsQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCubicFilterWeightsQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_FILTER_CUBIC_WEIGHTS_EXTENSION_NAME`
- `VK_QCOM_FILTER_CUBIC_WEIGHTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BLIT_IMAGE_CUBIC_WEIGHTS_INFO_QCOM`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUBIC_WEIGHTS_FEATURES_QCOM`
  - `VK_STRUCTURE_TYPE_SAMPLER_CUBIC_WEIGHTS_CREATE_INFO_QCOM`

## [](#_version_history)Version History

- Revision 1, 2023-06-23 (jleger)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_filter_cubic_weights)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0