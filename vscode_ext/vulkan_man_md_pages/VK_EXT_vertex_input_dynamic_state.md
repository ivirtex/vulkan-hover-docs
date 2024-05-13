# VK_EXT_vertex_input_dynamic_state(3) Manual Page

## Name

VK_EXT_vertex_input_dynamic_state - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

353

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_vertex_input_dynamic_state%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_vertex_input_dynamic_state%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-08-21

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Spencer Fricke, Samsung

- Stu Smith, AMD

## <a href="#_description" class="anchor"></a>Description

One of the states that contributes to the combinatorial explosion of
pipeline state objects that need to be created, is the vertex input
binding and attribute descriptions. By allowing them to be dynamic
applications may reduce the number of pipeline objects they need to
create.

This extension adds dynamic state support for what is normally static
state in
[VkPipelineVertexInputStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineVertexInputStateCreateInfo.html).

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription2EXT.html)

- [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVertexInputDynamicStateFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_VERTEX_INPUT_DYNAMIC_STATE_EXTENSION_NAME`

- `VK_EXT_VERTEX_INPUT_DYNAMIC_STATE_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_VERTEX_INPUT_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VERTEX_INPUT_DYNAMIC_STATE_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_ATTRIBUTE_DESCRIPTION_2_EXT`

  - `VK_STRUCTURE_TYPE_VERTEX_INPUT_BINDING_DESCRIPTION_2_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2020-11-05 (Piers Daniell)

  - Make
    [VkVertexInputBindingDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputBindingDescription2EXT.html)
    extensible

  - Add new
    [VkVertexInputAttributeDescription2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkVertexInputAttributeDescription2EXT.html)
    struct for the `pVertexAttributeDescriptions` parameter to
    [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html) so it is also
    extensible

- Revision 1, 2020-08-21 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_vertex_input_dynamic_state"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
