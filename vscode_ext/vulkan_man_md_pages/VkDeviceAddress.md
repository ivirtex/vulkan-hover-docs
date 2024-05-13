# VkDeviceAddress(3) Manual Page

## Name

VkDeviceAddress - Vulkan device address type



## <a href="#_c_specification" class="anchor"></a>C Specification

[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html) represents device buffer address
values:

``` c
// Provided by VK_VERSION_1_0
typedef uint64_t VkDeviceAddress;
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAccelerationStructureCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureCreateInfoKHR.html),
[VkBindIndexBufferIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindIndexBufferIndirectCommandNV.html),
[VkBindPipelineIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindPipelineIndirectCommandNV.html),
[VkBindVertexBufferIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindVertexBufferIndirectCommandNV.html),
[VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressCreateInfoEXT.html),
[VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html),
[VkCopyMemoryIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryIndirectCommandNV.html),
[VkCopyMemoryToImageIndirectCommandNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToImageIndirectCommandNV.html),
[VkDecompressMemoryRegionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDecompressMemoryRegionNV.html),
[VkDescriptorAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorAddressInfoEXT.html),
[VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingInfoEXT.html),
[VkDescriptorDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorDataEXT.html),
[VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingCallbackDataEXT.html),
[VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html),
[VkDeviceOrHostAddressConstAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstAMDX.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html),
[VkImageViewAddressPropertiesNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageViewAddressPropertiesNVX.html),
[VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapCreateInfoEXT.html),
[VkStridedDeviceAddressRegionKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStridedDeviceAddressRegionKHR.html),
[VkTraceRaysIndirectCommand2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkTraceRaysIndirectCommand2KHR.html),
[vkCmdBuildAccelerationStructuresIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildAccelerationStructuresIndirectKHR.html),
[vkCmdCopyMemoryIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryIndirectNV.html),
[vkCmdCopyMemoryToImageIndirectNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCopyMemoryToImageIndirectNV.html),
[vkCmdDecompressMemoryIndirectCountNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDecompressMemoryIndirectCountNV.html),
[vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphAMDX.html),
[vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectAMDX.html),
[vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdDispatchGraphIndirectCountAMDX.html),
[vkCmdInitializeGraphScratchMemoryAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdInitializeGraphScratchMemoryAMDX.html),
[vkCmdTraceRaysIndirect2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirect2KHR.html),
[vkCmdTraceRaysIndirectKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysIndirectKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceAddress"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
