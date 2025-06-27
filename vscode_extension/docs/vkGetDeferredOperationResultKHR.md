# vkGetDeferredOperationResultKHR(3) Manual Page

## Name

vkGetDeferredOperationResultKHR - Query the result of a deferred
operation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `vkGetDeferredOperationResultKHR` function is defined as:

``` c
// Provided by VK_KHR_deferred_host_operations
VkResult vkGetDeferredOperationResultKHR(
    VkDevice                                    device,
    VkDeferredOperationKHR                      operation);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device which owns `operation`.

- `operation` is the operation whose deferred result is being queried.

## <a href="#_description" class="anchor"></a>Description

If no command has been deferred on `operation`,
`vkGetDeferredOperationResultKHR` returns `VK_SUCCESS`.

If the deferred operation is pending, `vkGetDeferredOperationResultKHR`
returns `VK_NOT_READY`.

If the deferred operation is complete, it returns the appropriate return
value from the original command. This value **must** be one of the
[VkResult](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResult.html) values which could have been returned by the
original command if the operation had not been deferred.

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeferredOperationResultKHR-device-parameter"
  id="VUID-vkGetDeferredOperationResultKHR-device-parameter"></a>
  VUID-vkGetDeferredOperationResultKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeferredOperationResultKHR-operation-parameter"
  id="VUID-vkGetDeferredOperationResultKHR-operation-parameter"></a>
  VUID-vkGetDeferredOperationResultKHR-operation-parameter  
  `operation` **must** be a valid
  [VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html) handle

- <a href="#VUID-vkGetDeferredOperationResultKHR-operation-parent"
  id="VUID-vkGetDeferredOperationResultKHR-operation-parent"></a>
  VUID-vkGetDeferredOperationResultKHR-operation-parent  
  `operation` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

- `VK_NOT_READY`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_deferred_host_operations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_deferred_host_operations.html),
[VkDeferredOperationKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeferredOperationKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeferredOperationResultKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
