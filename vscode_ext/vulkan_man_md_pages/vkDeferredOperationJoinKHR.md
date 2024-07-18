# vkDeferredOperationJoinKHR(3) Manual Page

## Name

vkDeferredOperationJoinKHR - Assign a thread to a deferred operation



## <a href="#_c_specification" class="anchor"></a>C Specification

To assign a thread to a deferred operation, call:

``` c
// Provided by VK_KHR_deferred_host_operations
VkResult vkDeferredOperationJoinKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      operation);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `operation`.

- `operation` is the deferred operation that the calling thread should
  work on.

## <a href="#_description" class="anchor"></a>Description

The `vkDeferredOperationJoinKHR` command will execute a portion of the
deferred operation on the calling thread.

The return value will be one of the following:

- A return value of `VK_SUCCESS` indicates that `operation` is complete.
  The application **should** use
  [vkGetDeferredOperationResultKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeferredOperationResultKHR.html)
  to retrieve the result of `operation`.

- A return value of `VK_THREAD_DONE_KHR` indicates that the deferred
  operation is not complete, but there is no work remaining to assign to
  threads. Future calls to
  [vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html) are not
  necessary and will simply harm performance. This situation **may**
  occur when other threads executing
  [vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html) are
  about to complete `operation`, and the implementation is unable to
  partition the workload any further.

- A return value of `VK_THREAD_IDLE_KHR` indicates that the deferred
  operation is not complete, and there is no work for the thread to do
  at the time of the call. This situation **may** occur if the operation
  encounters a temporary reduction in parallelism. By returning
  `VK_THREAD_IDLE_KHR`, the implementation is signaling that it expects
  that more opportunities for parallelism will emerge as execution
  progresses, and that future calls to
  [vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html) **can**
  be beneficial. In the meantime, the application **can** perform other
  work on the calling thread.

Implementations **must** guarantee forward progress by enforcing the
following invariants:

1.  If only one thread has invoked
    [vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html) on a
    given operation, that thread **must** execute the operation to
    completion and return `VK_SUCCESS`.

2.  If multiple threads have concurrently invoked
    [vkDeferredOperationJoinKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDeferredOperationJoinKHR.html) on the
    same operation, then at least one of them **must** complete the
    operation and return `VK_SUCCESS`.

Valid Usage (Implicit)

- <a href="#VUID-vkDeferredOperationJoinKHR-device-parameter"
  id="VUID-vkDeferredOperationJoinKHR-device-parameter"></a>
  VUID-vkDeferredOperationJoinKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDeferredOperationJoinKHR-operation-parameter"
  id="VUID-vkDeferredOperationJoinKHR-operation-parameter"></a>
  VUID-vkDeferredOperationJoinKHR-operation-parameter  
  `operation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkDeferredOperationJoinKHR-operation-parent"
  id="VUID-vkDeferredOperationJoinKHR-operation-parent"></a>
  VUID-vkDeferredOperationJoinKHR-operation-parent  
  `operation` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_THREAD_DONE_KHR`

- `VK_THREAD_IDLE_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_deferred_host_operations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_deferred_host_operations.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDeferredOperationJoinKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
