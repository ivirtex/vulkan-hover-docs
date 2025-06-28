# VK\_AMD\_shader\_early\_and\_late\_fragment\_tests(3) Manual Page

## Name

VK\_AMD\_shader\_early\_and\_late\_fragment\_tests - device extension



## [](#_registered_extension_number)Registered Extension Number

322

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_get\_physical\_device\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_spir_v_dependencies)SPIR-V Dependencies

- [SPV\_AMD\_shader\_early\_and\_late\_fragment\_tests](https://github.khronos.org/SPIRV-Registry/extensions/AMD/SPV_AMD_shader_early_and_late_fragment_tests.html)

## [](#_contact)Contact

- Tobias Hector [\[GitHub\]tobski](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_early_and_late_fragment_tests%5D%20%40tobski%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_early_and_late_fragment_tests%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_AMD\_shader\_early\_and\_late\_fragment\_tests](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_AMD_shader_early_and_late_fragment_tests.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-09-14

**Interactions and External Dependencies**

- This extension interacts with `VK_EXT_shader_stencil_export`

**Contributors**

- Tobias Hector, AMD

## [](#_description)Description

This extension adds support for the [`SPV_AMD_shader_early_and_late_fragment_tests`](https://github.khronos.org/SPIRV-Registry/extensions/AMD/SPV_AMD_shader_early_and_late_fragment_tests.html) extension, allowing shaders to explicitly opt in to allowing both early *and* late fragment tests with the `EarlyAndLateFragmentTestsAMD` execution mode.

If the `VK_EXT_shader_stencil_export` extension is supported, additional execution modes allowing early depth tests similar to `DepthUnchanged`, `DepthLess`, and `DepthGreater` are provided.

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_AMD_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_EXTENSION_NAME`
- `VK_AMD_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_FEATURES_AMD`

## [](#_version_history)Version History

- Revision 1, 2021-09-14 (Tobias Hector)
  
  - Initial draft

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_AMD_shader_early_and_late_fragment_tests)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0