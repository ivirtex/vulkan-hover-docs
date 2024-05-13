# vkGetDeferredOperationMaxConcurrencyKHR(3) Manual Page

## Name

vkGetDeferredOperationMaxConcurrencyKHR - Query the maximum concurrency
on a deferred operation



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the number of additional threads that can usefully be joined to
a deferred operation, call:

``` c
// Provided by VK_KHR_deferred_host_operations
uint32_t vkGetDeferredOperationMaxConcurrencyKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      operation);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `operation`.

- `operation` is the deferred operation to be queried.

## <a href="#_description" class="anchor"></a>Description

The returned value is the maximum number of threads that can usefully
execute a deferred operation concurrently, reported for the state of the
deferred operation at the point this command is called. This value is
intended to be used to better schedule work onto available threads.
Applications **can** join any number of threads to the deferred
operation and expect it to eventually complete, though excessive joins
**may** return `VK_THREAD_DONE_KHR` immediately, performing no useful
work.

If `operation` is complete, `vkGetDeferredOperationMaxConcurrencyKHR`
returns zero.

If `operation` is currently joined to any threads, the value returned by
this command **may** immediately be out of date.

If `operation` is pending, implementations **must** not return zero
unless at least one thread is currently executing
[vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html) on
`operation`. If there are such threads, the implementation **should**
return an estimate of the number of additional threads which it could
profitably use.

Implementations **may** return 2<sup>32</sup>-1 to indicate that the
maximum concurrency is unknown and cannot be easily derived.
Implementations **may** return values larger than the maximum
concurrency available on the host CPU. In these situations, an
application **should** clamp the return value rather than
oversubscribing the machine.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>The recommended usage pattern for applications is to query this value
once, after deferral, and schedule no more than the specified number of
threads to join the operation. Each time a joined thread receives
<code>VK_THREAD_IDLE_KHR</code>, the application should schedule an
additional join at some point in the future, but is not required to do
so.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeferredOperationMaxConcurrencyKHR-device-parameter"
  id="VUID-vkGetDeferredOperationMaxConcurrencyKHR-device-parameter"></a>
  VUID-vkGetDeferredOperationMaxConcurrencyKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetDeferredOperationMaxConcurrencyKHR-operation-parameter"
  id="VUID-vkGetDeferredOperationMaxConcurrencyKHR-operation-parameter"></a>
  VUID-vkGetDeferredOperationMaxConcurrencyKHR-operation-parameter  
  `operation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkGetDeferredOperationMaxConcurrencyKHR-operation-parent"
  id="VUID-vkGetDeferredOperationMaxConcurrencyKHR-operation-parent"></a>
  VUID-vkGetDeferredOperationMaxConcurrencyKHR-operation-parent  
  `operation` **must** have been created, allocated, or retrieved from
  `device`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_deferred_host_operations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_deferred_host_operations.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeferredOperationMaxConcurrencyKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
