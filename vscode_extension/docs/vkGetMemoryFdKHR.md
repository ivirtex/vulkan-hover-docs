# vkGetMemoryFdKHR(3) Manual Page

## Name

vkGetMemoryFdKHR - Get a POSIX file descriptor for a memory object



## <a href="#_c_specification" class="anchor"></a>C Specification

To export a POSIX file descriptor referencing the payload of a Vulkan
device memory object, call:

``` c
// Provided by VK_KHR_external_memory_fd
VkResult vkGetMemoryFdKHR(
    VkDevice                                    device,
    const VkMemoryGetFdInfoKHR*                 pGetFdInfo,
    int*                                        pFd);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that created the device memory being
  exported.

- `pGetFdInfo` is a pointer to a
  [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetFdInfoKHR.html) structure containing
  parameters of the export operation.

- `pFd` will return a file descriptor referencing the payload of the
  device memory object.

## <a href="#_description" class="anchor"></a>Description

Each call to `vkGetMemoryFdKHR` **must** create a new file descriptor
holding a reference to the memory object’s payload and transfer
ownership of the file descriptor to the application. To avoid leaking
resources, the application **must** release ownership of the file
descriptor using the `close` system call when it is no longer needed, or
by importing a Vulkan memory object from it. Where supported by the
operating system, the implementation **must** set the file descriptor to
be closed automatically when an `execve` system call is made.

Valid Usage (Implicit)

- <a href="#VUID-vkGetMemoryFdKHR-device-parameter"
  id="VUID-vkGetMemoryFdKHR-device-parameter"></a>
  VUID-vkGetMemoryFdKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetMemoryFdKHR-pGetFdInfo-parameter"
  id="VUID-vkGetMemoryFdKHR-pGetFdInfo-parameter"></a>
  VUID-vkGetMemoryFdKHR-pGetFdInfo-parameter  
  `pGetFdInfo` **must** be a valid pointer to a valid
  [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetFdInfoKHR.html) structure

- <a href="#VUID-vkGetMemoryFdKHR-pFd-parameter"
  id="VUID-vkGetMemoryFdKHR-pFd-parameter"></a>
  VUID-vkGetMemoryFdKHR-pFd-parameter  
  `pFd` **must** be a valid pointer to an `int` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_TOO_MANY_OBJECTS`

- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_external_memory_fd](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory_fd.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetFdInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetMemoryFdKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
