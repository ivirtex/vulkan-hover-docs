# VK\_QCOM\_filter\_cubic\_clamp(3) Manual Page

## Name

VK\_QCOM\_filter\_cubic\_clamp - device extension



## [](#_registered_extension_number)Registered Extension Number

522

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_filter\_cubic](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_filter_cubic.html)  
and  
     [Vulkan Version 1.2](#versions-1.2)  
     or  
     [VK\_EXT\_sampler\_filter\_minmax](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sampler_filter_minmax.html)

## [](#_contact)Contact

- Matthew Netsch [\[GitHub\]mnetsch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_QCOM_filter_cubic_clamp%5D%20%40mnetsch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_QCOM_filter_cubic_clamp%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2023-08-02

**Contributors**

- Jeff Leger, Qualcomm Technologies, Inc.

## [](#_description)Description

This extension extends cubic filtering by adding the ability to enable an anti-ringing clamp. Cubic filtering samples from a 4x4 region of texels and computes a cubic weighted average of the region. In some cases, the resulting value is outside the range of any of the texels in the 4x4 region. This is sometimes referred to as “filter overshoot” or “filter ringing” and can occur when there is a sharp discontinuity in the 4x4 region being filtered. For some use cases this “ringing” can produces unacceptable artifacts.

The solution to the ringing problem is to clamp the post-cubic-filtered value to be within the max and min of texel values in the 4x4 region. While such “range clamping” can be performed in shader code, the additional texture fetches and clamping ALU operations can be costly.

Certain Adreno GPUs are able to perform the range clamp in the texture unit during cubic filtering at significant performance/power savings versus a shader-based clamping approach. This extension exposes such hardware functionality.

This extension extends [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html), adding `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM` which enables the range clamp operation.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCubicClampFeaturesQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCubicClampFeaturesQCOM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_QCOM_FILTER_CUBIC_CLAMP_EXTENSION_NAME`
- `VK_QCOM_FILTER_CUBIC_CLAMP_SPEC_VERSION`
- Extending [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html):
  
  - `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE_RANGECLAMP_QCOM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_CUBIC_CLAMP_FEATURES_QCOM`

## [](#_version_history)Version History

- Revision 1, 2023-08-02 (jleger)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_QCOM_filter_cubic_clamp).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0