# vkTrimCommandPool(3) Manual Page

## Name

vkTrimCommandPool - Trim a command pool



## <a href="#_c_specification" class="anchor"></a>C Specification

To trim a command pool, call:

``` c
// Provided by VK_VERSION_1_1
void vkTrimCommandPool(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    VkCommandPoolTrimFlags                      flags);
```

or the equivalent command

``` c
// Provided by VK_KHR_maintenance1
void vkTrimCommandPoolKHR(
    VkDevice                                    device,
    VkCommandPool                               commandPool,
    VkCommandPoolTrimFlags                      flags);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the command pool.

- `commandPool` is the command pool to trim.

- `flags` is reserved for future use.

## <a href="#_description" class="anchor"></a>Description

Trimming a command pool recycles unused memory from the command pool
back to the system. Command buffers allocated from the pool are not
affected by the command.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>This command provides applications with some control over the
internal memory allocations used by command pools.</p>
<p>Unused memory normally arises from command buffers that have been
recorded and later reset, such that they are no longer using the memory.
On reset, a command buffer can return memory to its command pool, but
the only way to release memory from a command pool to the system
requires calling <a
href="vkResetCommandPool.html">vkResetCommandPool</a>, which cannot be
executed while any command buffers from that pool are still in use.
Subsequent recording operations into command buffers will reuse this
memory but since total memory requirements fluctuate over time, unused
memory can accumulate.</p>
<p>In this situation, trimming a command pool <strong>may</strong> be
useful to return unused memory back to the system, returning the total
outstanding memory allocated by the pool back to a more “average”
value.</p>
<p>Implementations utilize many internal allocation strategies that make
it impossible to guarantee that all unused memory is released back to
the system. For instance, an implementation of a command pool
<strong>may</strong> involve allocating memory in bulk from the system
and sub-allocating from that memory. In such an implementation any live
command buffer that holds a reference to a bulk allocation would prevent
that allocation from being freed, even if only a small proportion of the
bulk allocation is in use.</p>
<p>In most cases trimming will result in a reduction in allocated but
unused memory, but it does not guarantee the “ideal” behavior.</p>
<p>Trimming <strong>may</strong> be an expensive operation, and
<strong>should</strong> not be called frequently. Trimming
<strong>should</strong> be treated as a way to relieve memory pressure
after application-known points when there exists enough unused memory
that the cost of trimming is “worth” it.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkTrimCommandPool-device-parameter"
  id="VUID-vkTrimCommandPool-device-parameter"></a>
  VUID-vkTrimCommandPool-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkTrimCommandPool-commandPool-parameter"
  id="VUID-vkTrimCommandPool-commandPool-parameter"></a>
  VUID-vkTrimCommandPool-commandPool-parameter  
  `commandPool` **must** be a valid [VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html)
  handle

- <a href="#VUID-vkTrimCommandPool-flags-zerobitmask"
  id="VUID-vkTrimCommandPool-flags-zerobitmask"></a>
  VUID-vkTrimCommandPool-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-vkTrimCommandPool-commandPool-parent"
  id="VUID-vkTrimCommandPool-commandPool-parent"></a>
  VUID-vkTrimCommandPool-commandPool-parent  
  `commandPool` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `commandPool` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkCommandPool](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPool.html),
[VkCommandPoolTrimFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandPoolTrimFlags.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkTrimCommandPool"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
