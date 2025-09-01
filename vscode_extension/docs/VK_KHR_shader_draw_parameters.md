# VK\_KHR\_shader\_draw\_parameters(3) Manual Page

## Name

VK\_KHR\_shader\_draw\_parameters - device extension



## [](#_registered_extension_number)Registered Extension Number

64

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_shader\_draw\_parameters](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_shader_draw_parameters.html)

## [](#_deprecation_state)Deprecation State

- *Promoted* to [Vulkan 1.1](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.1-promotions)

## [](#_contact)Contact

- Daniel Koch [\[GitHub\]dgkoch](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_draw_parameters%5D%20%40dgkoch%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_draw_parameters%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-09-05

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_shader_draw_parameters`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_draw_parameters.txt)

**Contributors**

- Daniel Koch, NVIDIA Corporation
- Jeff Bolz, NVIDIA
- Daniel Rakos, AMD
- Jan-Harald Fredriksen, ARM
- John Kessenich, Google
- Stuart Smith, IMG

## [](#_description)Description

This extension adds support for the following SPIR-V extension in Vulkan:

- `SPV_KHR_shader_draw_parameters`

The extension provides access to three additional built-in shader variables in Vulkan:

- `BaseInstance`, containing the `firstInstance` parameter passed to drawing commands,
- `BaseVertex`, containing the `firstVertex` or `vertexOffset` parameter passed to drawing commands, and
- `DrawIndex`, containing the index of the draw call currently being processed from an indirect drawing call.

When using GLSL source-based shader languages, the following variables from `GL_ARB_shader_draw_parameters` can map to these SPIR-V built-in decorations:

- `in int gl_BaseInstanceARB;` → `BaseInstance`,
- `in int gl_BaseVertexARB;` → `BaseVertex`, and
- `in int gl_DrawIDARB;` → `DrawIndex`.

## [](#_promotion_to_vulkan_1_1)Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1. However, the [`shaderDrawParameters`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderDrawParameters) feature bit was added to distinguish whether it is actually available or not.

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_DRAW_PARAMETERS_EXTENSION_NAME`
- `VK_KHR_SHADER_DRAW_PARAMETERS_SPEC_VERSION`

## [](#_new_built_in_variables)New Built-In Variables

- [`BaseInstance`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-baseinstance)
- [`BaseVertex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-basevertex)
- [`DrawIndex`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#interfaces-builtin-variables-drawindex)

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`DrawParameters`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-DrawParameters)

## [](#_issues)Issues

1\) Is this the same functionality as `GL_ARB_shader_draw_parameters`?

**RESOLVED**: It is actually a superset, as it also adds in support for arrayed drawing commands.

In GL for `GL_ARB_shader_draw_parameters`, `gl_BaseVertexARB` holds the integer value passed to the parameter to the command that resulted in the current shader invocation. In the case where the command has no `baseVertex` parameter, the value of `gl_BaseVertexARB` is zero. This means that `gl_BaseVertexARB` = `baseVertex` (for `glDrawElements` commands with `baseVertex`) or 0. In particular there are no `glDrawArrays` commands that take a `baseVertex` parameter.

Now in Vulkan, we have `BaseVertex` = `vertexOffset` (for indexed drawing commands) or `firstVertex` (for arrayed drawing commands), and so Vulkan’s version is really a superset of GL functionality.

## [](#_version_history)Version History

- Revision 1, 2016-10-05 (Daniel Koch)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_draw_parameters).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0