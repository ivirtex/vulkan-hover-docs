# VK\_KHR\_workgroup\_memory\_explicit\_layout(3) Manual Page

## Name

VK\_KHR\_workgroup\_memory\_explicit\_layout - device extension



## [](#_registered_extension_number)Registered Extension Number

337

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_workgroup\_memory\_explicit\_layout](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_workgroup_memory_explicit_layout.html)

## [](#_contact)Contact

- Caio Marcelo de Oliveira Filho [\[GitHub\]cmarcelo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_workgroup_memory_explicit_layout%5D%20%40cmarcelo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_workgroup_memory_explicit_layout%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2020-06-01

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- This extension provides API support for [`GL_EXT_shared_memory_block`](https://github.com/KhronosGroup/GLSL/blob/main/extensions/ext/GL_EXT_shared_memory_block.txt)

**Contributors**

- Caio Marcelo de Oliveira Filho, Intel
- Jeff Bolz, NVIDIA
- Graeme Leese, Broadcom
- Faith Ekstrand, Intel
- Daniel Koch, NVIDIA

## [](#_description)Description

This extension adds Vulkan support for the [`SPV_KHR_workgroup_memory_explicit_layout`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_workgroup_memory_explicit_layout.html) SPIR-V extension, which allows shaders to explicitly define the layout of `Workgroup` storage class memory and create aliases between variables from that storage class in a compute shader.

The aliasing feature allows different “views” on the same data, so the shader can bulk copy data from another storage class using one type (e.g. an array of large vectors), and then use the data with a more specific type. It also enables reducing the amount of workgroup memory consumed by allowing the shader to alias data whose lifetimes do not overlap.

The explicit layout support and some form of aliasing is also required for layering OpenCL on top of Vulkan.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceWorkgroupMemoryExplicitLayoutFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_EXTENSION_NAME`
- `VK_KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_FEATURES_KHR`

## [](#_new_spir_v_capabilities)New SPIR-V Capabilities

- [`WorkgroupMemoryExplicitLayoutKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-WorkgroupMemoryExplicitLayoutKHR)
- [`WorkgroupMemoryExplicitLayout8BitAccessKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-WorkgroupMemoryExplicitLayout8BitAccessKHR)
- [`WorkgroupMemoryExplicitLayout16BitAccessKHR`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#spirvenv-capabilities-table-WorkgroupMemoryExplicitLayout16BitAccessKHR)

## [](#_version_history)Version History

- Revision 1, 2020-06-01 (Caio Marcelo de Oliveira Filho)
  
  - Initial version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_workgroup_memory_explicit_layout).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0