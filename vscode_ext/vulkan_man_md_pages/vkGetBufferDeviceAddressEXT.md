# vkGetBufferDeviceAddress(3) Manual Page

## Name

vkGetBufferDeviceAddress - Query an address of a buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To query a 64-bit buffer device address value through which buffer
memory **can** be accessed in a shader, call:

``` c
// Provided by VK_VERSION_1_2
VkDeviceAddress vkGetBufferDeviceAddress(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_buffer_device_address
VkDeviceAddress vkGetBufferDeviceAddressKHR(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

or the equivalent command

``` c
// Provided by VK_EXT_buffer_device_address
VkDeviceAddress vkGetBufferDeviceAddressEXT(
    VkDevice                                    device,
    const VkBufferDeviceAddressInfo*            pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that the buffer was created on.

- `pInfo` is a pointer to a
  [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html) structure
  specifying the buffer to retrieve an address for.

## <a href="#_description" class="anchor"></a>Description

The 64-bit return value is an address of the start of `pInfo->buffer`.
The address range starting at this value and whose size is the size of
the buffer **can** be used in a shader to access the memory bound to
that buffer, using the `SPV_KHR_physical_storage_buffer` extension or
the equivalent `SPV_EXT_physical_storage_buffer` extension and the
`PhysicalStorageBuffer` storage class. For example, this value **can**
be stored in a uniform buffer, and the shader **can** read the value
from the uniform buffer and use it to do a dependent read/write to this
buffer. A value of zero is reserved as a “null” pointer and **must** not
be returned as a valid buffer device address. All loads, stores, and
atomics in a shader through `PhysicalStorageBuffer` pointers **must**
access addresses in the address range of some buffer.

If the buffer was created with a non-zero value of
[VkBufferOpaqueCaptureAddressCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferOpaqueCaptureAddressCreateInfo.html)::`opaqueCaptureAddress`
or
[VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressCreateInfoEXT.html)::`deviceAddress`,
the return value will be the same address that was returned at capture
time.

The returned address **must** satisfy the alignment requirement
specified by
[VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)::`alignment` for the
buffer in
[VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html)::`buffer`.

If multiple [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) objects are bound to overlapping
ranges of [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html), implementations **may**
return address ranges which overlap. In this case, it is ambiguous which
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) is associated with any given device address.
For purposes of valid usage, if multiple [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)
objects **can** be attributed to a device address, a
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) is selected such that valid usage passes, if
it exists.

Valid Usage

- <a href="#VUID-vkGetBufferDeviceAddress-bufferDeviceAddress-03324"
  id="VUID-vkGetBufferDeviceAddress-bufferDeviceAddress-03324"></a>
  VUID-vkGetBufferDeviceAddress-bufferDeviceAddress-03324  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddress"
  target="_blank" rel="noopener"><code>bufferDeviceAddress</code></a> or
  <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressEXT"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceBufferDeviceAddressFeaturesEXT</code>::<code>bufferDeviceAddress</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetBufferDeviceAddress-device-03325"
  id="VUID-vkGetBufferDeviceAddress-device-03325"></a>
  VUID-vkGetBufferDeviceAddress-device-03325  
  If `device` was created with multiple physical devices, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDevice"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressMultiDevice</code></a> or <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressMultiDeviceEXT"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceBufferDeviceAddressFeaturesEXT</code>::<code>bufferDeviceAddressMultiDevice</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetBufferDeviceAddress-device-parameter"
  id="VUID-vkGetBufferDeviceAddress-device-parameter"></a>
  VUID-vkGetBufferDeviceAddress-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetBufferDeviceAddress-pInfo-parameter"
  id="VUID-vkGetBufferDeviceAddress-pInfo-parameter"></a>
  VUID-vkGetBufferDeviceAddress-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkBufferDeviceAddressInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetBufferDeviceAddress"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
