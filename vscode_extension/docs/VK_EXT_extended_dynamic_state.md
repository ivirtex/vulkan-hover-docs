# VK\_EXT\_extended\_dynamic\_state(3) Manual Page

## Name

VK\_EXT\_extended\_dynamic\_state - device extension



## [](#_registered_extension_number)Registered Extension Number

268

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.3](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.3-promotions)

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_extended_dynamic_state%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_extended_dynamic_state%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-12-09

**IP Status**

No known IP claims.

**Contributors**

- Dan Ginsburg, Valve Corporation
- Graeme Leese, Broadcom
- Hans-Kristian Arntzen, Valve Corporation
- Jan-Harald Fredriksen, Arm Limited
- Faith Ekstrand, Intel
- Jeff Bolz, NVIDIA
- Jesse Hall, Google
- Philip Rebohle, Valve Corporation
- Stuart Smith, Imagination Technologies
- Tobias Hector, AMD

## [](#_description)Description

This extension adds some more dynamic state to support applications that need to reduce the number of pipeline state objects they compile and bind.

## [](#_new_commands)New Commands

- [vkCmdBindVertexBuffers2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers2EXT.html)
- [vkCmdSetCullModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetCullModeEXT.html)
- [vkCmdSetDepthBoundsTestEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthBoundsTestEnableEXT.html)
- [vkCmdSetDepthCompareOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthCompareOpEXT.html)
- [vkCmdSetDepthTestEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthTestEnableEXT.html)
- [vkCmdSetDepthWriteEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthWriteEnableEXT.html)
- [vkCmdSetFrontFaceEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetFrontFaceEXT.html)
- [vkCmdSetPrimitiveTopologyEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetPrimitiveTopologyEXT.html)
- [vkCmdSetScissorWithCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetScissorWithCountEXT.html)
- [vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilOpEXT.html)
- [vkCmdSetStencilTestEnableEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetStencilTestEnableEXT.html)
- [vkCmdSetViewportWithCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWithCountEXT.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceExtendedDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExtendedDynamicStateFeaturesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME`
- `VK_EXT_EXTENDED_DYNAMIC_STATE_SPEC_VERSION`
- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDynamicState.html):
  
  - `VK_DYNAMIC_STATE_CULL_MODE_EXT`
  - `VK_DYNAMIC_STATE_DEPTH_BOUNDS_TEST_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_DEPTH_COMPARE_OP_EXT`
  - `VK_DYNAMIC_STATE_DEPTH_TEST_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_DEPTH_WRITE_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_FRONT_FACE_EXT`
  - `VK_DYNAMIC_STATE_PRIMITIVE_TOPOLOGY_EXT`
  - `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT_EXT`
  - `VK_DYNAMIC_STATE_STENCIL_OP_EXT`
  - `VK_DYNAMIC_STATE_STENCIL_TEST_ENABLE_EXT`
  - `VK_DYNAMIC_STATE_VERTEX_INPUT_BINDING_STRIDE_EXT`
  - `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT`

## [](#_promotion_to_vulkan_1_3)Promotion to Vulkan 1.3

All dynamic state enumerants and commands in this extension are included in core Vulkan 1.3, with the EXT suffix omitted. The feature structure is not promoted. Extension interfaces that were promoted remain available as aliases of the core functionality.

## [](#_issues)Issues

1\) Why are the values of `pStrides` in [vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindVertexBuffers2.html) limited to be between 0 and the maximum extent of the binding, when this restriction is not present for the same static state?

Implementing these edge cases adds overhead to some implementations that would require significant cost when calling this function, and the intention is that this state should be more or less free to change.

[VK\_EXT\_vertex\_input\_dynamic\_state](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_vertex_input_dynamic_state.html) allows the stride to be changed freely when supported via [vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetVertexInputEXT.html).

## [](#_version_history)Version History

- Revision 1, 2019-12-09 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_extended_dynamic_state).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0