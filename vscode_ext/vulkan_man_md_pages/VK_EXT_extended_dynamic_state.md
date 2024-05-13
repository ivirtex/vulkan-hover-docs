# VK_EXT_extended_dynamic_state(3) Manual Page

## Name

VK_EXT_extended_dynamic_state - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

268

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.3-promotions"
  target="_blank" rel="noopener">Vulkan 1.3</a>

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_extended_dynamic_state%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_extended_dynamic_state%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

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

## <a href="#_description" class="anchor"></a>Description

This extension adds some more dynamic state to support applications that
need to reduce the number of pipeline state objects they compile and
bind.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBindVertexBuffers2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers2EXT.html)

- [vkCmdSetCullModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetCullModeEXT.html)

- [vkCmdSetDepthBoundsTestEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthBoundsTestEnableEXT.html)

- [vkCmdSetDepthCompareOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthCompareOpEXT.html)

- [vkCmdSetDepthTestEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthTestEnableEXT.html)

- [vkCmdSetDepthWriteEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDepthWriteEnableEXT.html)

- [vkCmdSetFrontFaceEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetFrontFaceEXT.html)

- [vkCmdSetPrimitiveTopologyEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetPrimitiveTopologyEXT.html)

- [vkCmdSetScissorWithCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetScissorWithCountEXT.html)

- [vkCmdSetStencilOpEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilOpEXT.html)

- [vkCmdSetStencilTestEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetStencilTestEnableEXT.html)

- [vkCmdSetViewportWithCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCountEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExtendedDynamicStateFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExtendedDynamicStateFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME`

- `VK_EXT_EXTENDED_DYNAMIC_STATE_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

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

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTENDED_DYNAMIC_STATE_FEATURES_EXT`

## <a href="#_promotion_to_vulkan_1_3" class="anchor"></a>Promotion to Vulkan 1.3

This extension has been partially promoted. All dynamic state enumerants
and entry points in this extension are included in core Vulkan 1.3, with
the EXT suffix omitted. The feature structure is not promoted. Extension
interfaces that were promoted remain available as aliases of the core
functionality.

## <a href="#_issues" class="anchor"></a>Issues

1\) Why are the values of `pStrides` in
[vkCmdBindVertexBuffers2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindVertexBuffers2.html) limited to be
between 0 and the maximum extent of the binding, when this restriction
is not present for the same static state?

Implementing these edge cases adds overhead to some implementations that
would require significant cost when calling this function, and the
intention is that this state should be more or less free to change.

[VK_EXT_vertex_input_dynamic_state](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_vertex_input_dynamic_state.html)
allows the stride to be changed freely when supported via
[vkCmdSetVertexInputEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetVertexInputEXT.html).

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-12-09 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_extended_dynamic_state"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
