# VK\_NV\_coverage\_reduction\_mode(3) Manual Page

## Name

VK\_NV\_coverage\_reduction\_mode - device extension



## [](#_registered_extension_number)Registered Extension Number

251

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_NV\_framebuffer\_mixed\_samples](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_framebuffer_mixed_samples.html)  
and  
     [VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Kedarnath Thangudu [\[GitHub\]kthangudu](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_coverage_reduction_mode%5D%20%40kthangudu%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_coverage_reduction_mode%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-01-29

**Contributors**

- Kedarnath Thangudu, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

When using a framebuffer with mixed samples, a per-fragment coverage reduction operation is performed which generates color sample coverage from the pixel coverage. This extension defines the following modes to control how this reduction is performed.

- Merge: When there are more samples in the pixel coverage than color samples, there is an implementation-dependent association of each pixel coverage sample to a color sample. In the merge mode, the color sample coverage is computed such that only if any associated sample in the pixel coverage is covered, the color sample is covered. This is the default mode.
- Truncate: When there are more raster samples (N) than color samples(M), there is one to one association of the first M raster samples to the M color samples; other raster samples are ignored.

When the number of raster samples is equal to the color samples, there is a one to one mapping between them in either of the above modes.

The new command [vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html) can be used to query the various raster, color, depth/stencil sample count and reduction mode combinations that are supported by the implementation. This extension would allow an implementation to support the behavior of both `VK_NV_framebuffer_mixed_samples` and `VK_AMD_mixed_attachment_samples` extensions simultaneously.

## [](#_new_commands)New Commands

- [vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html)

## [](#_new_structures)New Structures

- [VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebufferMixedSamplesCombinationNV.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceCoverageReductionModeFeaturesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceCoverageReductionModeFeaturesNV.html)
- Extending [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineMultisampleStateCreateInfo.html):
  
  - [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)

## [](#_new_enums)New Enums

- [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCoverageReductionModeNV.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineCoverageReductionStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCoverageReductionStateCreateFlagsNV.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_NV_COVERAGE_REDUCTION_MODE_EXTENSION_NAME`
- `VK_NV_COVERAGE_REDUCTION_MODE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV`
  - `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV`

## [](#_version_history)Version History

- Revision 1, 2019-01-29 (Kedarnath Thangudu)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_NV_coverage_reduction_mode)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0