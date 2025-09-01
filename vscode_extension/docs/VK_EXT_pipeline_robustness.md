# VK\_EXT\_pipeline\_robustness(3) Manual Page

## Name

VK\_EXT\_pipeline\_robustness - device extension



## [](#_registered_extension_number)Registered Extension Number

69

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.4](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.4-promotions)

## [](#_contact)Contact

- Jarred Davies

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2022-07-12

**Interactions and External Dependencies**

- Interacts with `VK_EXT_robustness2`
- Interacts with `VK_EXT_image_robustness`
- Interacts with `VK_KHR_ray_tracing_pipeline`

**Contributors**

- Jarred Davies, Imagination Technologies
- Alex Walters, Imagination Technologies
- Piers Daniell, NVIDIA
- Graeme Leese, Broadcom Corporation
- Jeff Leger, Qualcomm Technologies, Inc.
- Faith Ekstrand, Intel
- Lionel Landwerlin, Intel
- Shahbaz Youssefi, Google, Inc.

## [](#_description)Description

This extension allows users to request robustness on a per-pipeline stage basis.

As [`robustBufferAccess`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-robustBufferAccess) and other robustness features may have an adverse effect on performance, this extension is designed to allow users to request robust behavior only where it may be needed.

## [](#_new_structures)New Structures

- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html):
  
  - [VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessCreateInfoEXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePipelineRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineRobustnessFeaturesEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDevicePipelineRobustnessPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineRobustnessPropertiesEXT.html)

## [](#_new_enums)New Enums

- [VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehaviorEXT.html)
- [VkPipelineRobustnessImageBehaviorEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehaviorEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_PIPELINE_ROBUSTNESS_EXTENSION_NAME`
- `VK_EXT_PIPELINE_ROBUSTNESS_SPEC_VERSION`
- Extending [VkPipelineRobustnessBufferBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessBufferBehavior.html):
  
  - `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DEVICE_DEFAULT_EXT`
  - `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_DISABLED_EXT`
  - `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_2_EXT`
  - `VK_PIPELINE_ROBUSTNESS_BUFFER_BEHAVIOR_ROBUST_BUFFER_ACCESS_EXT`
- Extending [VkPipelineRobustnessImageBehavior](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRobustnessImageBehavior.html):
  
  - `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DEVICE_DEFAULT_EXT`
  - `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_DISABLED_EXT`
  - `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_2_EXT`
  - `VK_PIPELINE_ROBUSTNESS_IMAGE_BEHAVIOR_ROBUST_IMAGE_ACCESS_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PIPELINE_ROBUSTNESS_CREATE_INFO_EXT`

## [](#_promotion_to_vulkan_1_4)Promotion to Vulkan 1.4

Functionality in this extension is included in core Vulkan 1.4 with the EXT suffix omitted. The original type, enum and command names are still available as aliases of the core functionality.

## [](#_version_history)Version History

- Revision 1, 2022-07-12 (Jarred Davies)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_pipeline_robustness).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0