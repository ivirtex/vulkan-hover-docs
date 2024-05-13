# VK_EXT_pipeline_creation_feedback(3) Manual Page

## Name

VK_EXT_pipeline_creation_feedback - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

193

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_special_use" class="anchor"></a>Special Use

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#extendingvulkan-compatibility-specialuse"
  target="_blank" rel="noopener">Developer tools</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jean-Francois Roy <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_pipeline_creation_feedback%5D%20@jfroy%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_pipeline_creation_feedback%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jfroy</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds a mechanism to provide feedback to an application
about pipeline creation, with the specific goal of allowing a feedback
loop between build systems and in-the-field application executions to
ensure effective pipeline caches are shipped to customers.

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkPipelineCreationFeedbackEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackEXT.html)

- Extending
  [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html),
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html),
  [VkExecutionGraphPipelineCreateInfoAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExecutionGraphPipelineCreateInfoAMDX.html):

  - [VkPipelineCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackCreateInfoEXT.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkPipelineCreationFeedbackFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlagBitsEXT.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkPipelineCreationFeedbackFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineCreationFeedbackFlagsEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_PIPELINE_CREATION_FEEDBACK_EXTENSION_NAME`

- `VK_EXT_PIPELINE_CREATION_FEEDBACK_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PIPELINE_CREATION_FEEDBACK_CREATE_INFO_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

Functionality in this extension is included in core Vulkan 1.3, with the
EXT suffix omitted. The original type, enum and command names are still
available as aliases of the core functionality.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-03-12 (Jean-Francois Roy)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_pipeline_creation_feedback"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
