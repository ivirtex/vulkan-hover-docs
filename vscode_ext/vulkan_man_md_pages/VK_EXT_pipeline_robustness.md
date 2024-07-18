# VK_EXT_pipeline_robustness(3) Manual Page

## Name

VK_EXT_pipeline_robustness - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

69

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Jarred Davies

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-07-12

**Interactions and External Dependencies**  
- Interacts with [`VK_EXT_robustness2`](VK_EXT_robustness2.html)

- Interacts with
  [`VK_EXT_image_robustness`](VK_EXT_image_robustness.html)

- Interacts with
  [`VK_KHR_ray_tracing_pipeline`](VK_KHR_ray_tracing_pipeline.html)

**Contributors**  
- Jarred Davies, Imagination Technologies

- Alex Walters, Imagination Technologies

- Piers Daniell, NVIDIA

- Graeme Leese, Broadcom Corporation

- Jeff Leger, Qualcomm Technologies, Inc.

- Faith Ekstrand, Intel

- Lionel Landwerlin, Intel

- Shahbaz Youssefi, Google, Inc.

## <a href="#_description" class="anchor"></a>Description

This extension allows users to request robustness on a per-pipeline
stage basis.

As <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-robustBufferAccess"
target="_blank" rel="noopener"><code>robustBufferAccess</code></a> and
other robustness features may have an adverse effect on performance,
this extension is designed to allow users to request robust behavior
only where it may be needed.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
  [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html):

  - [VkPipelineRobustnessCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessCreateInfoEXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDevicePipelineRobustnessFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineRobustnessFeaturesEXT.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDevicePipelineRobustnessPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePipelineRobustnessPropertiesEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPipelineRobustnessBufferBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessBufferBehaviorEXT.html)

- [VkPipelineRobustnessImageBehaviorEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRobustnessImageBehaviorEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PIPELINE_ROBUSTNESS_EXTENSION_NAME`

- `VK_EXT_PIPELINE_ROBUSTNESS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_PIPELINE_ROBUSTNESS_PROPERTIES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_ROBUSTNESS_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-07-12 (Jarred Davies)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_pipeline_robustness"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
