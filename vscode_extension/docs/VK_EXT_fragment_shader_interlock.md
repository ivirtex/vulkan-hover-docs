# VK_EXT_fragment_shader_interlock(3) Manual Page

## Name

VK_EXT_fragment_shader_interlock - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

252

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_EXT_fragment_shader_interlock](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/EXT/SPV_EXT_fragment_shader_interlock.html)

## <a href="#_contact" class="anchor"></a>Contact

- Piers Daniell <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_fragment_shader_interlock%5D%20@pdaniell-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_fragment_shader_interlock%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>pdaniell-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2019-05-02

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_ARB_fragment_shader_interlock`](https://registry.khronos.org/OpenGL/extensions/ARB/ARB_fragment_shader_interlock.txt)

**Contributors**  
- Daniel Koch, NVIDIA

- Graeme Leese, Broadcom

- Jan-Harald Fredriksen, Arm

- Faith Ekstrand, Intel

- Jeff Bolz, NVIDIA

- Ruihao Zhang, Qualcomm

- Slawomir Grajewski, Intel

- Spencer Fricke, Samsung

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the `FragmentShaderPixelInterlockEXT`,
`FragmentShaderSampleInterlockEXT`, and
`FragmentShaderShadingRateInterlockEXT` capabilities from the
`SPV_EXT_fragment_shader_interlock` extension to Vulkan.

Enabling these capabilities provides a critical section for fragment
shaders to avoid overlapping pixels being processed at the same time,
and certain guarantees about the ordering of fragment shader invocations
of fragments of overlapping pixels.

This extension can be useful for algorithms that need to access
per-pixel data structures via shader loads and stores. Algorithms using
this extension can access per-pixel data structures in critical sections
without other invocations accessing the same per-pixel data.
Additionally, the ordering guarantees are useful for cases where the API
ordering of fragments is meaningful. For example, applications may be
able to execute programmable blending operations in the fragment shader,
where the destination buffer is read via image loads and the final value
is written via image stores.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFragmentShaderInterlockFeaturesEXT.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_EXT_FRAGMENT_SHADER_INTERLOCK_EXTENSION_NAME`

- `VK_EXT_FRAGMENT_SHADER_INTERLOCK_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_FRAGMENT_SHADER_INTERLOCK_FEATURES_EXT`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentShaderSampleInterlockEXT"
  target="_blank"
  rel="noopener"><code>FragmentShaderInterlockEXT</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentShaderPixelInterlockEXT"
  target="_blank"
  rel="noopener"><code>FragmentShaderPixelInterlockEXT</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-FragmentShaderShadingRateInterlockEXT"
  target="_blank"
  rel="noopener"><code>FragmentShaderShadingRateInterlockEXT</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2019-05-24 (Piers Daniell)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_EXT_fragment_shader_interlock"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
