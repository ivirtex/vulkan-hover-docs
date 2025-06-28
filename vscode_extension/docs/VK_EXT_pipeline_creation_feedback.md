# VK\_EXT\_pipeline\_creation\_feedback(3) Manual Page

## Name

VK\_EXT\_pipeline\_creation\_feedback - device extension



## [](#_registered_extension_number)Registered Extension Number

193

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_special_use)Special Use

- [Developer tools](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#extendingvulkan-compatibility-specialuse)

## [](#_contact)Contact

- Jean-Francois Roy [\[GitHub\]jfroy](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pipeline_creation_feedback%5D%20%40jfroy%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pipeline_creation_feedback%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-03-12

**IP Status**

No known IP claims.

**Contributors**

- Jean-Francois Roy, Google
- Hai Nguyen, Google
- Andrew Ellem, Google
- Bob Fraser, Google
- Sujeevan Rajayogam, Google
- Jan-Harald Fredriksen, ARM
- Jeff Leger, Qualcomm Technologies, Inc.
- Jeff Bolz, NVIDIA
- Daniel Koch, NVIDIA
- Neil Henning, AMD

## [](#_description)Description

This extension adds a mechanism to provide feedback to an application about pipeline creation, with the specific goal of allowing a feedback loop between build systems and in-the-field application executions to ensure effective pipeline caches are shipped to customers.

## [](#_new_structures)New Structures

- [VkPipelineCreationFeedbackEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackEXT.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html), [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html), [VkDataGraphPipelineCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineCreateInfoARM.html):
  
  - [VkPipelineCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackCreateInfoEXT.html)

## [](#_new_enums)New Enums

- [VkPipelineCreationFeedbackFlagBitsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlagBitsEXT.html)

## [](#_new_bitmasks)New Bitmasks

- [VkPipelineCreationFeedbackFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlagsEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PIPELINE_CREATION_FEEDBACK_EXTENSION_NAME`
- `VK_EXT_PIPELINE_CREATION_FEEDBACK_SPEC_VERSION`
- Extending [VkPipelineCreationFeedbackFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreationFeedbackFlagBits.html):
  
  - `VK_PIPELINE_CREATION_FEEDBACK_APPLICATION_PIPELINE_CACHE_HIT_BIT_EXT`
  - `VK_PIPELINE_CREATION_FEEDBACK_BASE_PIPELINE_ACCELERATION_BIT_EXT`
  - `VK_PIPELINE_CREATION_FEEDBACK_VALID_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

Vulkan APIs in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. External interactions defined by this extension, such as SPIR-V token names, retain their original names. The original Vulkan API names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2019-03-12 (Jean-Francois Roy)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_pipeline_creation_feedback)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0