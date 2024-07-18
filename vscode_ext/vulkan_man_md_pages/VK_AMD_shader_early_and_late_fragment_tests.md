# VK_AMD_shader_early_and_late_fragment_tests(3) Manual Page

## Name

VK_AMD_shader_early_and_late_fragment_tests - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

322

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_AMD_shader_early_and_late_fragment_tests](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/AMD/SPV_AMD_shader_early_and_late_fragment_tests.html)

## <a href="#_contact" class="anchor"></a>Contact

- Tobias Hector <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_AMD_shader_early_and_late_fragment_tests%5D%20@tobski%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_AMD_shader_early_and_late_fragment_tests%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>tobski</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_AMD_shader_early_and_late_fragment_tests](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_AMD_shader_early_and_late_fragment_tests.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-09-14

**Interactions and External Dependencies**  
- This extension interacts with
  [`VK_EXT_shader_stencil_export`](VK_EXT_shader_stencil_export.html)

**Contributors**  
- Tobias Hector, AMD

## <a href="#_description" class="anchor"></a>Description

This extension adds support for the
[`SPV_AMD_shader_early_and_late_fragment_tests`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/AMD/SPV_AMD_shader_early_and_late_fragment_tests.html)
extension, allowing shaders to explicitly opt in to allowing both early
*and* late fragment tests with the `EarlyAndLateFragmentTestsAMD`
execution mode.

If [`VK_EXT_shader_stencil_export`](VK_EXT_shader_stencil_export.html)
is supported, additional execution modes allowing early depth tests
similar to `DepthUnchanged`, `DepthLess`, and `DepthGreater` are
provided.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderEarlyAndLateFragmentTestsFeaturesAMD.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_AMD_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_EXTENSION_NAME`

- `VK_AMD_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_EARLY_AND_LATE_FRAGMENT_TESTS_FEATURES_AMD`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-09-14 (Tobias Hector)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_AMD_shader_early_and_late_fragment_tests"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
