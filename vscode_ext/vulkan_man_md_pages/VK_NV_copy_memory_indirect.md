# VK_NV_copy_memory_indirect(3) Manual Page

## Name

VK_NV_copy_memory_indirect - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

427

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

        
[VK_KHR_get_physical_device_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_physical_device_properties2.html)  
         or  
         [Version 1.1](#versions-1.1)  
     and  
     [VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html)  
or  
[Version 1.2](#versions-1.2)  

## <a href="#_contact" class="anchor"></a>Contact

- Vikram Kushwaha <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NV_copy_memory_indirect%5D%20@vkushwaha-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NV_copy_memory_indirect%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>vkushwaha-nv</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2022-10-14

**Contributors**  
- Vikram Kushwaha, NVIDIA

- Jeff Bolz, NVIDIA

- Christoph Kubisch, NVIDIA

- Daniel Koch, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension adds support for performing copies between memory and
image regions using indirect parameters that are read by the device from
a buffer during execution. This functionality **may** be useful for
performing copies where the copy parameters are not known during the
command buffer creation time.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdCopyMemoryIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryIndirectNV.html)

- [vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToImageIndirectNV.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCopyMemoryIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryIndirectCommandNV.html)

- [VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageIndirectCommandNV.html)

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceCopyMemoryIndirectFeaturesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCopyMemoryIndirectFeaturesNV.html)

- Extending
  [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html):

  - [VkPhysicalDeviceCopyMemoryIndirectPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceCopyMemoryIndirectPropertiesNV.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NV_COPY_MEMORY_INDIRECT_EXTENSION_NAME`

- `VK_NV_COPY_MEMORY_INDIRECT_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_FEATURES_NV`

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_COPY_MEMORY_INDIRECT_PROPERTIES_NV`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2022-10-14 (Vikram Kushwaha)

  - Initial draft

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NV_copy_memory_indirect"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
