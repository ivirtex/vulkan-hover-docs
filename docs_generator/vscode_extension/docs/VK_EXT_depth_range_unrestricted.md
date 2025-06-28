# VK\_EXT\_depth\_range\_unrestricted(3) Manual Page

## Name

VK\_EXT\_depth\_range\_unrestricted - device extension



## [](#_registered_extension_number)Registered Extension Number

14

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_contact)Contact

- Piers Daniell [\[GitHub\]pdaniell-nv](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_depth_range_unrestricted%5D%20%40pdaniell-nv%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_depth_range_unrestricted%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-06-22

**Contributors**

- Daniel Koch, NVIDIA
- Jeff Bolz, NVIDIA

## [](#_description)Description

This extension removes the [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) `minDepth` and `maxDepth` restrictions that the values must be between `0.0` and `1.0`, inclusive. It also removes the same restriction on [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineDepthStencilStateCreateInfo.html) `minDepthBounds` and `maxDepthBounds`. Finally it removes the restriction on the `depth` value in [VkClearDepthStencilValue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClearDepthStencilValue.html).

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_DEPTH_RANGE_UNRESTRICTED_EXTENSION_NAME`
- `VK_EXT_DEPTH_RANGE_UNRESTRICTED_SPEC_VERSION`

## [](#_issues)Issues

1\) How do [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) `minDepth` and `maxDepth` values outside of the `0.0` to `1.0` range interact with [Primitive Clipping](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-clipping)?

**RESOLVED**: The behavior described in [Primitive Clipping](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-clipping) still applies. If depth clamping is disabled the depth values are still clipped to 0 ≤ zc ≤ wc before the viewport transform. If depth clamping is enabled the above equation is ignored and the depth values are instead clamped to the [VkViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/VkViewport.html) `minDepth` and `maxDepth` values, which in the case of this extension can be outside of the `0.0` to `1.0` range.

2\) What happens if a resulting depth fragment is outside of the `0.0` to `1.0` range and the depth buffer is fixed-point rather than floating-point?

**RESOLVED**: This situation can also arise without this extension (when fragment shaders replace depth values, for example), and this extension does not change the behavior, which is defined in the [Depth Test](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragops-depth) section of the Fragment Operations chapter.

## [](#_version_history)Version History

- Revision 1, 2017-06-22 (Piers Daniell)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_depth_range_unrestricted)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0