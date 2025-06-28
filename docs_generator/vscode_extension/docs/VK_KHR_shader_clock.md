# VK\_KHR\_shader\_clock(3) Manual Page

## Name

VK\_KHR\_shader\_clock - device extension



## [](#_registered_extension_number)Registered Extension Number

182

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_shader\_clock](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_shader_clock.html)

## [](#_contact)Contact

- Aaron Hagan [\[GitHub\]ahagan](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_clock%5D%20%40ahagan%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_clock%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-4-25

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_ARB_shader_clock`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_shader_clock.txt) and [`GL_EXT_shader_realtime_clock`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shader_realtime_clock.txt)

**Contributors**

- Aaron Hagan, AMD
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension advertises the SPIR-V `ShaderClockKHR` capability for Vulkan, which allows a shader to query a real-time or monotonically incrementing counter at the subgroup level or across the device level. The two valid SPIR-V scopes for `OpReadClockKHR` are `Subgroup` and `Device`.

When using GLSL source-based shading languages, the `clockRealtime*EXT`() timing functions map to the `OpReadClockKHR` instruction with a scope of `Device`, and the `clock*ARB`() timing functions map to the `OpReadClockKHR` instruction with a scope of `Subgroup`.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderClockFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderClockFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_CLOCK_EXTENSION_NAME`
- `VK_KHR_SHADER_CLOCK_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_CLOCK_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`ShaderClockKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-ShaderClockKHR)

## [](#_version_history)Version History

- Revision 1, 2019-4-25 (Aaron Hagan)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_clock)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0