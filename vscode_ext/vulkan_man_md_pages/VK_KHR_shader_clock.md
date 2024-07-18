# VK_KHR_shader_clock(3) Manual Page

## Name

VK_KHR_shader_clock - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

182

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_shader_clock](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_shader_clock.html)

## <a href="#_contact" class="anchor"></a>Contact

- Aaron Hagan <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_clock%5D%20@ahagan%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_clock%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ahagan</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-4-25

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_shader_clock`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_clock.txt)
  and
  [`GL_EXT_shader_realtime_clock`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shader_realtime_clock.txt)

**Contributors**  
- Aaron Hagan, AMD

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension advertises the SPIR-V `ShaderClockKHR` capability for
Vulkan, which allows a shader to query a real-time or monotonically
incrementing counter at the subgroup level or across the device level.
The two valid SPIR-V scopes for `OpReadClockKHR` are `Subgroup` and
`Device`.

When using GLSL source-based shading languages, the
`clockRealtime*EXT`() timing functions map to the `OpReadClockKHR`
instruction with a scope of `Device`, and the `clock*ARB`() timing
functions map to the `OpReadClockKHR` instruction with a scope of
`Subgroup`.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderClockFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderClockFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_CLOCK_EXTENSION_NAME`

- `VK_KHR_SHADER_CLOCK_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CLOCK_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-ShaderClockKHR"
  target="_blank" rel="noopener"><code>ShaderClockKHR</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-4-25 (Aaron Hagan)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_clock"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
