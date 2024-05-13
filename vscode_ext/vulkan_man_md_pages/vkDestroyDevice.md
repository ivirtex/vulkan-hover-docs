# vkDestroyDevice(3) Manual Page

## Name

vkDestroyDevice - Destroy a logical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To destroy a device, call:

``` c
// Provided by VK_VERSION_1_0
void vkDestroyDevice(
    VkDevice                                    device,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device to destroy.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

To ensure that no work is active on the device,
[vkDeviceWaitIdle](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeviceWaitIdle.html) **can** be used to gate the
destruction of the device. Prior to destroying a device, an application
is responsible for destroying/freeing any Vulkan objects that were
created using that device as the first parameter of the corresponding
`vkCreate*` or `vkAllocate*` command.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The lifetime of each of these objects is bound by the lifetime of the
<code>VkDevice</code> object. Therefore, to avoid resource leaks, it is
critical that an application explicitly free all of these resources
prior to calling <code>vkDestroyDevice</code>.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkDestroyDevice-device-05137"
  id="VUID-vkDestroyDevice-device-05137"></a>
  VUID-vkDestroyDevice-device-05137  
  All child objects created on `device` **must** have been destroyed
  prior to destroying `device`

- <a href="#VUID-vkDestroyDevice-device-00379"
  id="VUID-vkDestroyDevice-device-00379"></a>
  VUID-vkDestroyDevice-device-00379  
  If `VkAllocationCallbacks` were provided when `device` was created, a
  compatible set of callbacks **must** be provided here

- <a href="#VUID-vkDestroyDevice-device-00380"
  id="VUID-vkDestroyDevice-device-00380"></a>
  VUID-vkDestroyDevice-device-00380  
  If no `VkAllocationCallbacks` were provided when `device` was created,
  `pAllocator` **must** be `NULL`

Valid Usage (Implicit)

- <a href="#VUID-vkDestroyDevice-device-parameter"
  id="VUID-vkDestroyDevice-device-parameter"></a>
  VUID-vkDestroyDevice-device-parameter  
  If `device` is not `NULL`, `device` **must** be a valid
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDestroyDevice-pAllocator-parameter"
  id="VUID-vkDestroyDevice-pAllocator-parameter"></a>
  VUID-vkDestroyDevice-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

Host Synchronization

- Host access to `device` **must** be externally synchronized

- Host access to all `VkQueue` objects created from `device` **must** be
  externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDestroyDevice"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
