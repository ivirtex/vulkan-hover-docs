# VK_KHR_shader_atomic_int64(3) Manual Page

## Name

VK_KHR_shader_atomic_int64 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

181

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Aaron Hagan <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_atomic_int64%5D%20@ahagan%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_atomic_int64%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ahagan</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-07-05

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_gpu_shader_int64`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_gpu_shader_int64.txt)
  and
  [`GL_EXT_shader_atomic_int64`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shader_atomic_int64.txt)

**Contributors**  
- Aaron Hagan, AMD

- Daniel Rakos, AMD

- Jeff Bolz, NVIDIA

- Neil Henning, Codeplay

## <a href="#_description" class="anchor"></a>Description

This extension advertises the SPIR-V **Int64Atomics** capability for
Vulkan, which allows a shader to contain 64-bit atomic operations on
signed and unsigned integers. The supported operations include
OpAtomicMin, OpAtomicMax, OpAtomicAnd, OpAtomicOr, OpAtomicXor,
OpAtomicAdd, OpAtomicExchange, and OpAtomicCompareExchange.

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. However, if Vulkan 1.2 is supported and this
extension is not, the `shaderBufferInt64Atomics` capability is optional.
The original type, enum and command names are still available as aliases
of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderAtomicInt64FeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderAtomicInt64FeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_ATOMIC_INT64_EXTENSION_NAME`

- `VK_KHR_SHADER_ATOMIC_INT64_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_ATOMIC_INT64_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-Int64Atomics"
  target="_blank" rel="noopener"><code>Int64Atomics</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-07-05 (Aaron Hagan)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_atomic_int64"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
