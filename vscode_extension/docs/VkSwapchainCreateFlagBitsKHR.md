# VkSwapchainCreateFlagBitsKHR(3) Manual Page

## Name

VkSwapchainCreateFlagBitsKHR - Bitmask controlling swapchain creation



## [](#_c_specification)C Specification

Bits which **can** be set in [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html)::`flags`, specifying parameters of swapchain creation, are:

```c++
// Provided by VK_KHR_swapchain
typedef enum VkSwapchainCreateFlagBitsKHR {
  // Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
    VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR = 0x00000001,
  // Provided by VK_VERSION_1_1 with VK_KHR_swapchain
    VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR = 0x00000002,
  // Provided by VK_KHR_swapchain_mutable_format
    VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR = 0x00000004,
  // Provided by VK_EXT_swapchain_maintenance1
    VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT = 0x00000008,
  // Provided by VK_KHR_present_id2
    VK_SWAPCHAIN_CREATE_PRESENT_ID_2_BIT_KHR = 0x00000040,
  // Provided by VK_KHR_present_wait2
    VK_SWAPCHAIN_CREATE_PRESENT_WAIT_2_BIT_KHR = 0x00000080,
} VkSwapchainCreateFlagBitsKHR;
```

## [](#_description)Description

- `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR` specifies that images created from the swapchain (i.e. with the `swapchain` member of [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageSwapchainCreateInfoKHR.html) set to this swapchain’s handle) **must** use `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT`.
- `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` specifies that images created from the swapchain are protected images.
- `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` specifies that the images of the swapchain **can** be used to create a `VkImageView` with a different format than what the swapchain was created with. The list of allowed image view formats is specified by adding a [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatListCreateInfo.html) structure to the `pNext` chain of [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateInfoKHR.html). In addition, this flag also specifies that the swapchain **can** be created with usage flags that are not supported for the format the swapchain is created with but are supported for at least one of the allowed image view formats.
- `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT` specifies that the implementation **may** defer allocation of memory associated with each swapchain image until its index is to be returned from [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html) or [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html) for the first time.
- `VK_SWAPCHAIN_CREATE_PRESENT_ID_2_BIT_KHR` specifies that applications **can** include the `VkPresentId2KHR` structure in the `pNext` chain of the [VkPresentInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentInfoKHR.html) structure to associate an identifier with each presentation request.
- `VK_SWAPCHAIN_CREATE_PRESENT_WAIT_2_BIT_KHR` specifies that applications **can** use `vkWaitForPresent2KHR` to wait for the presentation engine to have begun presentation of the presentation request associated with [VkPresentWait2InfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPresentWait2InfoKHR.html)::`presentId` on `swapchain`.

## [](#_see_also)See Also

[VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VkSwapchainCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainCreateFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSwapchainCreateFlagBitsKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0