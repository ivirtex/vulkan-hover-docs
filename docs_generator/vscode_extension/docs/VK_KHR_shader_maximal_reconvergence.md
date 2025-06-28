# VK\_KHR\_shader\_maximal\_reconvergence(3) Manual Page

## Name

VK\_KHR\_shader\_maximal\_reconvergence - device extension



## [](#_registered_extension_number)Registered Extension Number

435

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_KHR\_maximal\_reconvergence](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_maximal_reconvergence.html)

## [](#_contact)Contact

- Alan Baker [\[GitHub\]alan-baker](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_maximal_reconvergence%5D%20%40alan-baker%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_maximal_reconvergence%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_KHR\_shader\_maximal\_reconvergence](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_maximal_reconvergence.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-11-12

**IP Status**

No known IP claims.

**Interactions and External Dependencies**

- Requires SPIR-V 1.3.
- This extension requires [`SPV_KHR_maximal_reconvergence`](https://github.khronos.org/SPIRV-Registry/extensions/KHR/SPV_KHR_maximal_reconvergence.html)

**Contributors**

- Alan Baker, Google

## [](#_description)Description

This extension allows the use of the `SPV_KHR_maximal_reconvergence` SPIR-V extension in shader modules. `SPV_KHR_maximal_reconvergence` provides stronger guarantees that diverged subgroups will reconverge. These guarantees should match shader author intuition about divergence and reconvergence of invocations based on the structure of the code in the HLL.

Developers should utilize this extension if they require stronger guarantees about reconvergence than either the core spec or SPV\_KHR\_subgroup\_uniform\_control\_flow. This extension will define the rules that govern how invocations diverge and reconverge in a way that should match developer intuition. It allows robust programs to be written relying on subgroup operations and other tangled instructions.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_SHADER_MAXIMAL_RECONVERGENCE_EXTENSION_NAME`
- `VK_KHR_SHADER_MAXIMAL_RECONVERGENCE_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MAXIMAL_RECONVERGENCE_FEATURES_KHR`

## [](#_version_history)Version History

- Revision 1, 2021-11-12 (Alan Baker)
  
  - Internal draft version

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_shader_maximal_reconvergence)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0