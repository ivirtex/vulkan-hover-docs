# vkDestroyOpticalFlowSessionNV(3) Manual Page

## Name

vkDestroyOpticalFlowSessionNV - Destroy optical flow session object



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy an optical flow session object, call:

``` c
// Provided by VK_NV_optical_flow
void vkDestroyOpticalFlowSessionNV(
    VkDevice                                    device,
    VkOpticalFlowSessionNV                      session,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that was used for the creation of the optical
  flow session.

- `session` is the optical flow session to be destroyed.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyOpticalFlowSessionNV-device-parameter"
  id="VUID-vkDestroyOpticalFlowSessionNV-device-parameter"></a>
  VUID-vkDestroyOpticalFlowSessionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyOpticalFlowSessionNV-session-parameter"
  id="VUID-vkDestroyOpticalFlowSessionNV-session-parameter"></a>
  VUID-vkDestroyOpticalFlowSessionNV-session-parameter  
  `session` **must** be a valid
  [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html) handle

- <a href="#VUID-vkDestroyOpticalFlowSessionNV-pAllocator-parameter"
  id="VUID-vkDestroyOpticalFlowSessionNV-pAllocator-parameter"></a>
  VUID-vkDestroyOpticalFlowSessionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkDestroyOpticalFlowSessionNV-session-parent"
  id="VUID-vkDestroyOpticalFlowSessionNV-session-parent"></a>
  VUID-vkDestroyOpticalFlowSessionNV-session-parent  
  `session` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyOpticalFlowSessionNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
