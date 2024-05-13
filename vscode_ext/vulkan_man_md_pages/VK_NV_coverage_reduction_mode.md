# VK_NV_coverage_reduction_mode(3) Manual Page

## Name

VK_NV_coverage_reduction_mode - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

251

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_NV_framebuffer_mixed_samples](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_framebuffer_mixed_samples.html)  
and  
    
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
     or  
     [Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Kedarnath Thangudu <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_coverage_reduction_mode%5D%20@kthangudu%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_coverage_reduction_mode%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>kthangudu</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-01-29

**Contributors**  
- Kedarnath Thangudu, NVIDIA

- Jeff Bolz, NVIDIA

## <a href="#_description" class="anchor"></a>Description

When using a framebuffer with mixed samples, a per-fragment coverage
reduction operation is performed which generates color sample coverage
from the pixel coverage. This extension defines the following modes to
control how this reduction is performed.

- Merge: When there are more samples in the pixel coverage than color
  samples, there is an implementation-dependent association of each
  pixel coverage sample to a color sample. In the merge mode, the color
  sample coverage is computed such that only if any associated sample in
  the pixel coverage is covered, the color sample is covered. This is
  the default mode.

- Truncate: When there are more raster samples (N) than color
  samples(M), there is one to one association of the first M raster
  samples to the M color samples; other raster samples are ignored.

When the number of raster samples is equal to the color samples, there
is a one to one mapping between them in either of the above modes.

The new command
[vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html)
can be used to query the various raster, color, depth/stencil sample
count and reduction mode combinations that are supported by the
implementation. This extension would allow an implementation to support
the behavior of both `VK_NV_framebuffer_mixed_samples` and
`VK_AMD_mixed_attachment_samples` extensions simultaneously.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSupportedFramebufferMixedSamplesCombinationsNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkFramebufferMixedSamplesCombinationNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebufferMixedSamplesCombinationNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCoverageReductionModeFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCoverageReductionModeFeaturesNV.html)

- Extending
  [VkPipelineMultisampleStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineMultisampleStateCreateInfo.html):

  - [VkPipelineCoverageReductionStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateInfoNV.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkCoverageReductionModeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoverageReductionModeNV.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPipelineCoverageReductionStateCreateFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCoverageReductionStateCreateFlagsNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_COVERAGE_REDUCTION_MODE_EXTENSION_NAME`

- `VK_NV_COVERAGE_REDUCTION_MODE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_FRAMEBUFFER_MIXED_SAMPLES_COMBINATION_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COVERAGE_REDUCTION_MODE_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PIPELINE_COVERAGE_REDUCTION_STATE_CREATE_INFO_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-01-29 (Kedarnath Thangudu)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_coverage_reduction_mode"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
