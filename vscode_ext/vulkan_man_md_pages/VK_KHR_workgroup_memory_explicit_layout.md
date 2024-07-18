# VK_KHR_workgroup_memory_explicit_layout(3) Manual Page

## Name

VK_KHR_workgroup_memory_explicit_layout - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

337

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_workgroup_memory_explicit_layout](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_workgroup_memory_explicit_layout.html)

## <a href="#_contact" class="anchor"></a>Contact

- Caio Marcelo de Oliveira Filho <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_workgroup_memory_explicit_layout%5D%20@cmarcelo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_workgroup_memory_explicit_layout%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cmarcelo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2020-06-01

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- This extension provides API support for
  [`GL_EXT_shared_memory_block`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shared_memory_block.txt)

**Contributors**  
- Caio Marcelo de Oliveira Filho, Intel

- Jeff Bolz, NVIDIA

- Graeme Leese, Broadcom

- Faith Ekstrand, Intel

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds Vulkan support for the
[`SPV_KHR_workgroup_memory_explicit_layout`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_workgroup_memory_explicit_layout.html)
SPIR-V extension, which allows shaders to explicitly define the layout
of `Workgroup` storage class memory and create aliases between variables
from that storage class in a compute shader.

The aliasing feature allows different “views” on the same data, so the
shader can bulk copy data from another storage class using one type
(e.g. an array of large vectors), and then use the data with a more
specific type. It also enables reducing the amount of workgroup memory
consumed by allowing the shader to alias data whose lifetimes do not
overlap.

The explicit layout support and some form of aliasing is also required
for layering OpenCL on top of Vulkan.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_EXTENSION_NAME`

- `VK_KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-WorkgroupMemoryExplicitLayoutKHR"
  target="_blank"
  rel="noopener"><code>WorkgroupMemoryExplicitLayoutKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-WorkgroupMemoryExplicitLayout8BitAccessKHR"
  target="_blank"
  rel="noopener"><code>WorkgroupMemoryExplicitLayout8BitAccessKHR</code></a>

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-WorkgroupMemoryExplicitLayout16BitAccessKHR"
  target="_blank"
  rel="noopener"><code>WorkgroupMemoryExplicitLayout16BitAccessKHR</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2020-06-01 (Caio Marcelo de Oliveira Filho)

  - Initial version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_workgroup_memory_explicit_layout"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
