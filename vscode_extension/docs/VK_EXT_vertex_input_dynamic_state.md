# VK\_EXT\_vertex\_input\_dynamic\_state(3) Manual Page

## Name

VK\_EXT\_vertex\_input\_dynamic\_state - device extension



## [](#_registered_extension_number)Registered Extension Number

353

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_vertex_input_dynamic_state%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_vertex_input_dynamic_state%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-08-21

**IP Status**

No known IP claims.

**Contributors**

- Jeff Bolz, NVIDIA
- Spencer Fricke, Samsung
- Stu Smith, AMD

## [](#_description)Description

One of the states that contributes to the combinatorial explosion of pipeline state objects that need to be created, is the vertex input binding and attribute descriptions. By allowing them to be dynamic applications may reduce the number of pipeline objects they need to create.

This extension adds dynamic state support for what is normally static state in [VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineVertexInputStateCreateInfo.html).

## [](#_new_commands)New Commands

- [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html)

## [](#_new_structures)New Structures

- [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html)
- [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html)
- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_VERTEX_INPUT_DYNAMIC_STATE_EXTENSION_NAME`
- `VK_EXT_VERTEX_INPUT_DYNAMIC_STATE_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_INPUT_DYNAMIC_STATE_FEATURES_EXT`
  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT`
  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT`

## [](#_version_history)Version History

- Revision 2, 2020-11-05 (Piers Daniell)
  
  - Make [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputBindingDescription2EXT.html) extensible
  - Add new [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkVertexInputAttributeDescription2EXT.html) structure for the `pVertexAttributeDescriptions` parameter to [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html) so it is also extensible
- Revision 1, 2020-08-21 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_vertex_input_dynamic_state).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0