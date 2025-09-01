# VK\_ARM\_pipeline\_opacity\_micromap(3) Manual Page

## Name

VK\_ARM\_pipeline\_opacity\_micromap - device extension



## [](#_registered_extension_number)Registered Extension Number

597

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html)

## [](#_contact)Contact

- Mathieu Robart [\[GitHub\]mathieurobart-arm](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_ARM_pipeline_opacity_micromap%5D%20%40mathieurobart-arm%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_ARM_pipeline_opacity_micromap%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_ARM\_pipeline\_opacity\_micromap](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_ARM_pipeline_opacity_micromap.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2025-01-07

**IP Status**

No known IP claims.

**Contributors**

- Mathieu Robart, Arm
- Marius Bjorge, Arm
- Yaozhong Zhang, Arm
- Jan-Harald Fredriksen, Arm

## [](#_description)Description

The Opacity Micromap extension `VK_EXT_opacity_micromap` supports the new pipeline creation flag `VK_PIPELINE_CREATE_RAY_TRACING_OPACITY_MICROMAP_BIT_EXT`, indicating that the ray tracing pipeline may be used with acceleration structures referencing micromaps. This allows for possible optimizations, knowing beforehand that opacity micromaps may be used with the pipeline.

An equivalent flag does not exist for pipelines supporting Ray Query with opacity micromaps, such as graphics and compute. Consequently, it is currently not possible to optimize such pipelines for no-opacity, e.g. when opacity micromaps are supported by an application but not used by the pipeline. This may lead to performance degradation.

This extension adds a new flag, `VK_PIPELINE_CREATE_2_DISALLOW_OPACITY_MICROMAP_BIT_ARM`, indicating that a pipeline will NOT be used with an acceleration structure referencing an opacity micromap, therefore allowing possible pipeline optimizations.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePipelineOpacityMicromapFeaturesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineOpacityMicromapFeaturesARM.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_ARM_PIPELINE_OPACITY_MICROMAP_EXTENSION_NAME`
- `VK_ARM_PIPELINE_OPACITY_MICROMAP_SPEC_VERSION`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_DISALLOW_OPACITY_MICROMAP_BIT_ARM`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_OPACITY_MICROMAP_FEATURES_ARM`

## [](#_issues)Issues

None.

## [](#_examples)Examples

None.

## [](#_version_history)Version History

- Revision 1, 2025-01-07 (Mathieu Robart)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_ARM_pipeline_opacity_micromap).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0