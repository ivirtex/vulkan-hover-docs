# VkDeviceGroupSwapchainCreateInfoKHR(3) Manual Page

## Name

VkDeviceGroupSwapchainCreateInfoKHR - Structure specifying parameters of
a newly created swapchain object



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) includes a
`VkDeviceGroupSwapchainCreateInfoKHR` structure, then that structure
includes a set of device group present modes that the swapchain **can**
be used with.

The `VkDeviceGroupSwapchainCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkDeviceGroupSwapchainCreateInfoKHR {
    VkStructureType                     sType;
    const void*                         pNext;
    VkDeviceGroupPresentModeFlagsKHR    modes;
} VkDeviceGroupSwapchainCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `modes` is a bitfield of modes that the swapchain **can** be used
  with.

## <a href="#_description" class="anchor"></a>Description

If this structure is not present, `modes` is considered to be
`VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR`.

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceGroupSwapchainCreateInfoKHR-sType-sType"
  id="VUID-VkDeviceGroupSwapchainCreateInfoKHR-sType-sType"></a>
  VUID-VkDeviceGroupSwapchainCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEVICE_GROUP_SWAPCHAIN_CREATE_INFO_KHR`

- <a href="#VUID-VkDeviceGroupSwapchainCreateInfoKHR-modes-parameter"
  id="VUID-VkDeviceGroupSwapchainCreateInfoKHR-modes-parameter"></a>
  VUID-VkDeviceGroupSwapchainCreateInfoKHR-modes-parameter  
  `modes` **must** be a valid combination of
  [VkDeviceGroupPresentModeFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagBitsKHR.html)
  values

- <a
  href="#VUID-VkDeviceGroupSwapchainCreateInfoKHR-modes-requiredbitmask"
  id="VUID-VkDeviceGroupSwapchainCreateInfoKHR-modes-requiredbitmask"></a>
  VUID-VkDeviceGroupSwapchainCreateInfoKHR-modes-requiredbitmask  
  `modes` **must** not be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceGroupSwapchainCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
