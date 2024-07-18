# VK_NV_scissor_exclusive(3) Manual Page

## Name

VK_NV_scissor_exclusive - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

206

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Pat Brown <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_scissor_exclusive%5D%20@nvpbrown%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_scissor_exclusive%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>nvpbrown</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2023-01-18

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
None

**Contributors**  
- Pat Brown, NVIDIA

- Jeff Bolz, NVIDIA

- Piers Daniell, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for an exclusive scissor test to Vulkan. The
exclusive scissor test behaves like the scissor test, except that the
exclusive scissor test fails for pixels inside the corresponding
rectangle and passes for pixels outside the rectangle. If the same
rectangle is used for both the scissor and exclusive scissor tests, the
exclusive scissor test will pass if and only if the scissor test fails.

Version 2 of this extension introduces
`VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV` and
[vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorEnableNV.html).
Applications that use this dynamic state must ensure the implementation
advertises at least `specVersion` `2` of this extension.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetExclusiveScissorEnableNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorEnableNV.html)

- [vkCmdSetExclusiveScissorNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetExclusiveScissorNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceExclusiveScissorFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExclusiveScissorFeaturesNV.html)

- Extending
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html):

  - [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_SCISSOR_EXCLUSIVE_EXTENSION_NAME`

- `VK_NV_SCISSOR_EXCLUSIVE_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_ENABLE_NV`

  - `VK_DYNAMIC_STATE_EXCLUSIVE_SCISSOR_NV`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXCLUSIVE_SCISSOR_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_EXCLUSIVE_SCISSOR_STATE_CREATE_INFO_NV`

## <a href="#_issues" class="anchor"></a>Issues

1\) For the scissor test, the viewport state must be created with a
matching number of scissor and viewport rectangles. Should we have the
same requirement for exclusive scissors?

**RESOLVED**: For exclusive scissors, we relax this requirement and
allow an exclusive scissor rectangle count that is either zero or equal
to the number of viewport rectangles. If you pass in an exclusive
scissor count of zero, the exclusive scissor test is treated as
disabled.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2023-01-18 (Piers Daniell)

  - Add dynamic state for explicit exclusive scissor enables

- Revision 1, 2018-07-31 (Pat Brown)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_scissor_exclusive"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
