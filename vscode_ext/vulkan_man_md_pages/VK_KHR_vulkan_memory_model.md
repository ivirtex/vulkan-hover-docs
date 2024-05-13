# VK_KHR_vulkan_memory_model(3) Manual Page

## Name

VK_KHR_vulkan_memory_model - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

212

## <a href="#_revision" class="anchor"></a>Revision

3

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_spir_v_dependencies" class="anchor"></a>SPIR-V Dependencies

- [SPV_KHR_vulkan_memory_model](https://htmlpreview.github.io/?https://github.com/KhronosGroup/SPIRV-Registry/blob/main/extensions/KHR/SPV_KHR_vulkan_memory_model.html)

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.2-promotions"
  target="_blank" rel="noopener">Vulkan 1.2</a>

## <a href="#_contact" class="anchor"></a>Contact

- Jeff Bolz <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_vulkan_memory_model%5D%20@jeffbolznv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_vulkan_memory_model%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>jeffbolznv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2018-12-10

**IP Status**  
No known IP claims.

**Contributors**  
- Jeff Bolz, NVIDIA

- Alan Baker, Google

- Tobias Hector, AMD

- David Neto, Google

- Robert Simpson, Qualcomm Technologies, Inc.

- Brian Sumner, AMD

## <a href="#_description" class="anchor"></a>Description

The [VK_KHR_vulkan_memory_model](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_vulkan_memory_model.html)
extension allows use of the features guarded by the `VulkanMemoryModel`,
`VulkanMemoryModelDeviceScope`, and
`VulkanMemoryModelAvailabilityVisibilityChains` capabilities in shader
modules. The <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-model"
target="_blank" rel="noopener">Vulkan Memory Model</a> formally defines
how to synchronize memory accesses to the same memory locations
performed by multiple shader invocations.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Version 3 of the spec added a member
(<code>vulkanMemoryModelAvailabilityVisibilityChains</code>) to <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeaturesKHR.html">VkPhysicalDeviceVulkanMemoryModelFeaturesKHR</a>,
which is an incompatible change from version 2.</p></td>
</tr>
</tbody>
</table>

## <a href="#_promotion_to_vulkan_1_2" class="anchor"></a>Promotion to Vulkan 1.2

All functionality in this extension is included in core Vulkan 1.2, with
the KHR suffix omitted. However, if Vulkan 1.2 is supported and this
extension is not, the `vulkanMemoryModel` capability is optional. The
original type, enum and command names are still available as aliases of
the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceVulkanMemoryModelFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceVulkanMemoryModelFeaturesKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_VULKAN_MEMORY_MODEL_EXTENSION_NAME`

- `VK_KHR_VULKAN_MEMORY_MODEL_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_VULKAN_MEMORY_MODEL_FEATURES_KHR`

## <a href="#_new_spir_v_capabilities" class="anchor"></a>New SPIR-V Capabilities

- <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#spirvenv-capabilities-table-VulkanMemoryModel"
  target="_blank" rel="noopener"><code>VulkanMemoryModelKHR</code></a>

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2018-06-24 (Jeff Bolz)

  - Initial draft

- Revision 3, 2018-12-10 (Jeff Bolz)

  - Add vulkanMemoryModelAvailabilityVisibilityChains member to the
    VkPhysicalDeviceVulkanMemoryModelFeaturesKHR structure.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_vulkan_memory_model"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
