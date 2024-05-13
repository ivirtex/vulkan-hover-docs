# vkCreateOpticalFlowSessionNV(3) Manual Page

## Name

vkCreateOpticalFlowSessionNV - Creates an optical flow session object



## <a href="#_c_specification" class="anchor"></a>C Specification

To create an optical flow session object, call:

``` c
// Provided by VK_NV_optical_flow
VkResult vkCreateOpticalFlowSessionNV(
    VkDevice                                    device,
    const VkOpticalFlowSessionCreateInfoNV*     pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkOpticalFlowSessionNV*                     pSession);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the optical flow session
  object.

- `pCreateInfo` is a pointer to a
  [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html)
  structure containing parameters specifying the creation of the optical
  flow session.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pSession` is a pointer to a
  [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html) handle
  specifying the optical flow session object which will be created by
  this function when it returns `VK_SUCCESS`

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateOpticalFlowSessionNV-device-parameter"
  id="VUID-vkCreateOpticalFlowSessionNV-device-parameter"></a>
  VUID-vkCreateOpticalFlowSessionNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateOpticalFlowSessionNV-pCreateInfo-parameter"
  id="VUID-vkCreateOpticalFlowSessionNV-pCreateInfo-parameter"></a>
  VUID-vkCreateOpticalFlowSessionNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html)
  structure

- <a href="#VUID-vkCreateOpticalFlowSessionNV-pAllocator-parameter"
  id="VUID-vkCreateOpticalFlowSessionNV-pAllocator-parameter"></a>
  VUID-vkCreateOpticalFlowSessionNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateOpticalFlowSessionNV-pSession-parameter"
  id="VUID-vkCreateOpticalFlowSessionNV-pSession-parameter"></a>
  VUID-vkCreateOpticalFlowSessionNV-pSession-parameter  
  `pSession` **must** be a valid pointer to a
  [VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_optical_flow](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_optical_flow.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkOpticalFlowSessionCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionCreateInfoNV.html),
[VkOpticalFlowSessionNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpticalFlowSessionNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateOpticalFlowSessionNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
