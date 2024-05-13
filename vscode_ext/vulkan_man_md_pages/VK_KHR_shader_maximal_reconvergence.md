# VK_KHR_shader_maximal_reconvergence(3) Manual Page

## Name

VK_KHR_shader_maximal_reconvergence - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

435

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_maximal_reconvergence](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_maximal_reconvergence.html)

## <a href="#_contact" class="anchor"></a>Contact

- Alan Baker <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_shader_maximal_reconvergence%5D%20@alan-baker%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_shader_maximal_reconvergence%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>alan-baker</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_KHR_shader_maximal_reconvergence](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_KHR_shader_maximal_reconvergence.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-11-12

**IP Status**  
No known IP claims.

**Interactions and External Dependencies**  
- Requires SPIR-V 1.3.

- This extension requires
  [`SPV_KHR_maximal_reconvergence`](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_maximal_reconvergence.html)

**Contributors**  
- Alan Baker, Google

## <a href="#_description" class="anchor"></a>Description

This extension allows the use of the `SPV_KHR_maximal_reconvergence`
SPIR-V extension in shader modules. `SPV_KHR_maximal_reconvergence`
provides stronger guarantees that diverged subgroups will reconverge.
These guarantees should match shader author intuition about divergence
and reconvergence of invocations based on the structure of the code in
the HLL.

Developers should utilize this extension if they require stronger
guarantees about reconvergence than either the core spec or
SPV_KHR_subgroup_uniform_control_flow. This extension will define the
rules that govern how invocations diverge and reconverge in a way that
should match developer intuition. It allows robust programs to be
written relying on subgroup operations and other tangled instructions.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceShaderMaximalReconvergenceFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_SHADER_MAXIMAL_RECONVERGENCE_EXTENSION_NAME`

- `VK_KHR_SHADER_MAXIMAL_RECONVERGENCE_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADER_MAXIMAL_RECONVERGENCE_FEATURES_KHR`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-11-12 (Alan Baker)

  - Internal draft version

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_shader_maximal_reconvergence"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
