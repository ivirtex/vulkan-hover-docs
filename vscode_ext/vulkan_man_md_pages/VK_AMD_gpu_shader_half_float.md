# VK_AMD_gpu_shader_half_float(3) Manual Page

## Name

VK_AMD_gpu_shader_half_float - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

37

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_AMD_gpu_shader_half_float](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/AMD/SPV_AMD_gpu_shader_half_float.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Deprecated* by
  [VK_KHR_shader_float16_int8](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float16_int8.html)
  extension

  - Which in turn was *promoted* to <a
    href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
    target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Dominik Witczak <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_gpu_shader_half_float%5D%20@dominikwitczakamd%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_gpu_shader_half_float%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>dominikwitczakamd</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-04-11

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_AMD_gpu_shader_half_float`](https://registry.khronos.org/OpenGL/extensions/AMD/AMD_gpu_shader_half_float.txt)

**Contributors**  
- Daniel Rakos, AMD

- Dominik Witczak, AMD

- Donglin Wei, AMD

- Graham Sellers, AMD

- Qun Lin, AMD

- Rex Xu, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds support for using half float variables in shaders.

## <a href="#_deprecation_by_vk_khr_shader_float16_int8"
class="anchor"></a>Deprecation by `VK_KHR_shader_float16_int8`

Functionality in this extension was included in
[`VK_KHR_shader_float16_int8`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shader_float16_int8.html)
extension, when
[VkPhysicalDeviceShaderFloat16Int8FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderFloat16Int8FeaturesKHR.html)::`shaderFloat16`
is enabled.

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_GPU_SHADER_HALF_FLOAT_EXTENSION_NAME`

- `VK_AMD_GPU_SHADER_HALF_FLOAT_SPEC_VERSION`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2019-04-11 (Tobias Hector)

  - Marked as deprecated

- Revision 1, 2016-09-21 (Dominik Witczak)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_gpu_shader_half_float"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
