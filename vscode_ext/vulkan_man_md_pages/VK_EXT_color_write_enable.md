# VK_EXT_color_write_enable(3) Manual Page

## Name

VK_EXT_color_write_enable - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

382

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- Sharif Elcott <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_color_write_enable%5D%20@selcott%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_color_write_enable%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>selcott</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-02-25

**IP Status**  
No known IP claims.

**Contributors**  
- Sharif Elcott, Google

- Tobias Hector, AMD

- Piers Daniell, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows for selectively enabling and disabling writes to
output color attachments via a pipeline dynamic state.

The intended use cases for this new state are mostly identical to those
of colorWriteMask, such as selectively disabling writes to avoid
feedback loops between subpasses or bandwidth savings for unused
outputs. By making the state dynamic, one additional benefit is the
ability to reduce pipeline counts and pipeline switching via shaders
that write a superset of the desired data of which subsets are selected
dynamically. The reason for a new state, colorWriteEnable, rather than
making colorWriteMask dynamic is that, on many implementations, the more
flexible per-component semantics of the colorWriteMask state cannot be
made dynamic in a performant manner.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdSetColorWriteEnableEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetColorWriteEnableEXT.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceColorWriteEnableFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceColorWriteEnableFeaturesEXT.html)

- Extending
  [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorBlendStateCreateInfo.html):

  - [VkPipelineColorWriteCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineColorWriteCreateInfoEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_COLOR_WRITE_ENABLE_EXTENSION_NAME`

- `VK_EXT_COLOR_WRITE_ENABLE_SPEC_VERSION`

- Extending [VkDynamicState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDynamicState.html):

  - `VK_DYNAMIC_STATE_COLOR_WRITE_ENABLE_EXT`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COLOR_WRITE_ENABLE_FEATURES_EXT`

  - `VK_STRUCTURE_TYPE_PIPELINE_COLOR_WRITE_CREATE_INFO_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-01-25 (Sharif Elcott)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_color_write_enable"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
