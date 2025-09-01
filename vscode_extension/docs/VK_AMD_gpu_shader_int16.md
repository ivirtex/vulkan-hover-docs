# VK\_AMD\_gpu\_shader\_int16(3) Manual Page

## Name

VK\_AMD\_gpu\_shader\_int16 - device extension



## [](#_registered_extension_number)Registered Extension Number

133

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

None

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_AMD\_gpu\_shader\_int16](https://github.khronos.org/SPIRV-Registry/extensions/AMD/SPV_AMD_gpu_shader_int16.html)

## [](#_deprecation_state)Deprecation State

- *Deprecated* by [VK\_KHR\_shader\_float16\_int8](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_shader_float16_int8.html) extension
  
  - Which in turn was *promoted* to [Vulkan 1.2](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#versions-1.2-promotions)

## [](#_contact)Contact

- Qun Lin [\[GitHub\]linqun](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_gpu_shader_int16%5D%20%40linqun%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_gpu_shader_int16%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2019-04-11

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_AMD_gpu_shader_int16`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_gpu_shader_int16.txt)

**Contributors**

- Daniel Rakos, AMD
- Dominik Witczak, AMD
- Matthaeus G. Chajdas, AMD
- Rex Xu, AMD
- Timothy Lottes, AMD
- Zhi Cai, AMD

## [](#_description)Description

This extension adds support for using 16-bit integer variables in shaders.

## [](#_deprecation_by_vk_khr_shader_float16_int8)Deprecation by `VK_KHR_shader_float16_int8`

Functionality in this extension is included in the `VK_KHR_shader_float16_int8` extension, when the [`shaderInt16`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderInt16) and [`shaderFloat16`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-shaderFloat16) features are enabled.

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_GPU_SHADER_INT16_EXTENSION_NAME`
- `VK_AMD_GPU_SHADER_INT16_SPEC_VERSION`

## [](#_version_history)Version History

- Revision 2, 2019-04-11 (Tobias Hector)
  
  - Marked as deprecated
- Revision 1, 2017-06-18 (Dominik Witczak)
  
  - First version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_gpu_shader_int16).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0