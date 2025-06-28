# VK\_KHR\_pipeline\_binary(3) Manual Page

## Name

VK\_KHR\_pipeline\_binary - device extension



## [](#_registered_extension_number)Registered Extension Number

484

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html)  
or  
[Vulkan Version 1.4](#versions-1.4)

## [](#_contact)Contact

- Stu Smith [\[GitHub\]stu-s](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_pipeline_binary%5D%20%40stu-s%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_pipeline_binary%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_pipeline\_binary](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_pipeline_binary.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2024-07-01

**Contributors**

- Stu Smith, AMD
- Tobias Hector, AMD
- Alan Harrison, AMD
- Maciej Jesionowski, AMD
- Younggwan Kim, Arm
- Jan-Harald Fredriksen, Arm
- Ting Wei, Arm
- Chris Glover, Google
- Shahbaz Youssefi, Google
- Jakub Kuderski, Google
- Piotr Byszewski, Mobica
- Piers Daniell, NVIDIA
- Ralph Potter, Samsung
- Matthew Netsch, Qualcomm
- Hans-Kristian Arntzen, Valve
- Samuel Pitoiset, Valve
- Tatsuyuki Ishi, Valve

## [](#_description)Description

This extension provides a method to obtain binary data associated with individual pipelines such that applications can manage caching themselves instead of using VkPipelineCache objects.

## [](#_new_object_types)New Object Types

- [VkPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKHR.html)

## [](#_new_commands)New Commands

- [vkCreatePipelineBinariesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineBinariesKHR.html)
- [vkDestroyPipelineBinaryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPipelineBinaryKHR.html)
- [vkGetPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineBinaryDataKHR.html)
- [vkGetPipelineKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineKeyKHR.html)
- [vkReleaseCapturedPipelineDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkReleaseCapturedPipelineDataKHR.html)

## [](#_new_structures)New Structures

- [VkPipelineBinaryCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryCreateInfoKHR.html)
- [VkPipelineBinaryDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataInfoKHR.html)
- [VkPipelineBinaryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryDataKHR.html)
- [VkPipelineBinaryHandlesInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryHandlesInfoKHR.html)
- [VkPipelineBinaryKeyKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeyKHR.html)
- [VkPipelineBinaryKeysAndDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryKeysAndDataKHR.html)
- [VkPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateInfoKHR.html)
- [VkReleaseCapturedPipelineDataInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkReleaseCapturedPipelineDataInfoKHR.html)
- Extending [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkDevicePipelineBinaryInternalCacheControlKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevicePipelineBinaryInternalCacheControlKHR.html)
- Extending [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html), [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html), [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html):
  
  - [VkPipelineBinaryInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineBinaryInfoKHR.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDevicePipelineBinaryFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineBinaryFeaturesKHR.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDevicePipelineBinaryPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevicePipelineBinaryPropertiesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_PIPELINE_BINARY_EXTENSION_NAME`
- `VK_KHR_PIPELINE_BINARY_SPEC_VERSION`
- `VK_MAX_PIPELINE_BINARY_KEY_SIZE_KHR`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_PIPELINE_BINARY_KHR`
- Extending [VkPipelineCreateFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineCreateFlagBits2.html):
  
  - `VK_PIPELINE_CREATE_2_CAPTURE_DATA_BIT_KHR`
- Extending [VkResult](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResult.html):
  
  - `VK_ERROR_NOT_ENOUGH_SPACE_KHR`
  - `VK_PIPELINE_BINARY_MISSING_KHR`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_DEVICE_PIPELINE_BINARY_INTERNAL_CACHE_CONTROL_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_BINARY_FEATURES_KHR`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_BINARY_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_BINARY_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_BINARY_DATA_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_BINARY_HANDLES_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_BINARY_INFO_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_BINARY_KEY_KHR`
  - `VK_STRUCTURE_TYPE_PIPELINE_CREATE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_RELEASE_CAPTURED_PIPELINE_DATA_INFO_KHR`

## [](#_version_history)Version History

- Revision 1, 2021-12-10 (Chris Glover)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_pipeline_binary)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0