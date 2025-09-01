# VK\_EXT\_subgroup\_size\_control(3) Manual Page

## Name

VK\_EXT\_subgroup\_size\_control - device extension



## [](#_registered_extension_number)Registered Extension Number

226

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Neil Henning [\[GitHub\]sheredom](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_subgroup_size_control%5D%20%40sheredom%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_subgroup_size_control%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-03-05

**Contributors**

- Jeff Bolz, NVIDIA
- Faith Ekstrand, Intel
- Sławek Grajewski, Intel
- Jesse Hall, Google
- Neil Henning, AMD
- Daniel Koch, NVIDIA
- Jeff Leger, Qualcomm
- Graeme Leese, Broadcom
- Allan MacKinnon, Google
- Mariusz Merecki, Intel
- Graham Wihlidal, Electronic Arts

## [](#_description)Description

This extension enables an implementation to control the subgroup size by allowing a varying subgroup size and also specifying a required subgroup size.

It extends the subgroup support in Vulkan 1.1 to allow an implementation to expose a varying subgroup size. Previously Vulkan exposed a single subgroup size per physical device, with the expectation that implementations will behave as if all subgroups have the same size. Some implementations **may** dispatch shaders with a varying subgroup size for different subgroups. As a result they could implicitly split a large subgroup into smaller subgroups or represent a small subgroup as a larger subgroup, some of whose invocations were inactive on launch.

To aid developers in understanding the performance characteristics of their programs, this extension exposes a minimum and maximum subgroup size that a physical device supports and a pipeline create flag to enable that pipeline to vary its subgroup size. If enabled, any `SubgroupSize` decorated variables in the SPIR-V shader modules provided to pipeline creation **may** vary between the [minimum](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minSubgroupSize) and [maximum](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubgroupSize) subgroup sizes.

An implementation is also optionally allowed to support specifying a required subgroup size for a given pipeline stage. Implementations advertise which [stages support a required subgroup size](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-requiredSubgroupSizeStages), and any pipeline of a supported stage can be passed a [VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT.html) structure to set the subgroup size for that shader stage of the pipeline. For compute shaders, this requires the developer to query the [`maxComputeWorkgroupSubgroups`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxComputeWorkgroupSubgroups) and ensure that:

s=WorkGroupSize.x×WorkGroupSize.y×WorkgroupSize.z≤SubgroupSize×maxComputeWorkgroupSubgroups

Developers can also specify a new pipeline shader stage create flag that requires the implementation to have fully populated subgroups within local workgroups. This requires the workgroup size in the X dimension to be a multiple of the subgroup size.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceSubgroupSizeControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupSizeControlFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceSubgroupSizeControlPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupSizeControlPropertiesEXT.html)
- Extending [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html):
  
  - [VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_SUBGROUP_SIZE_CONTROL_EXTENSION_NAME`
- `VK_EXT_SUBGROUP_SIZE_CONTROL_SPEC_VERSION`
- Extending [VkPipelineShaderStageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateFlagBits.html):
  
  - `VK_PIPELINE_SHADER_STAGE_CREATE_ALLOW_VARYING_SUBGROUP_SIZE_BIT_EXT`
  - `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SUBGROUP_SIZE_CONTROL_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2019-03-05 (Neil Henning)
  
  - Initial draft
- Revision 2, 2019-07-26 (Faith Ekstrand)
  
  - Add the missing [VkPhysicalDeviceSubgroupSizeControlFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceSubgroupSizeControlFeaturesEXT.html) for querying subgroup size control features.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_subgroup_size_control).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0