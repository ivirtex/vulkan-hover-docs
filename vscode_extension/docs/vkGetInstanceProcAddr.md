# vkGetInstanceProcAddr(3) Manual Page

## Name

vkGetInstanceProcAddr - Return a function pointer for a command



## [](#_c_specification)C Specification

Function pointers for all Vulkan commands **can** be obtained by calling:

```c++
// Provided by VK_VERSION_1_0
PFN_vkVoidFunction vkGetInstanceProcAddr(
    VkInstance                                  instance,
    const char*                                 pName);
```

## [](#_parameters)Parameters

- `instance` is the instance that the function pointer will be compatible with, or `NULL` for commands not dependent on any instance.
- `pName` is the name of the command to obtain.

## [](#_description)Description

`vkGetInstanceProcAddr` itself is obtained in a platform- and loader- specific manner. Typically, the loader library will export this command as a function symbol, so applications **can** link against the loader library, or load it dynamically and look up the symbol using platform-specific APIs.

The table below defines the various use cases for `vkGetInstanceProcAddr` and expected return value (“fp” is “function pointer”) for each case. A valid returned function pointer (“fp”) **must** not be `NULL`.

The returned function pointer is of type [PFN\_vkVoidFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkVoidFunction.html), and **must** be cast to the type of the command being queried before use.

Table 1. `vkGetInstanceProcAddr` behavior    `instance` `pName` return value

\*1

`NULL`

undefined

invalid non-`NULL` instance

\*1

undefined

`NULL`

*global command*2

fp

`NULL`

[vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html)

fp5

instance

[vkGetInstanceProcAddr](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetInstanceProcAddr.html)

fp

instance

core *dispatchable command*

fp3

instance

enabled instance extension dispatchable command for `instance`

fp3

instance

available device extension4 dispatchable command for `instance`

fp3

any other case, not covered above

`NULL`

1

"\*" means any representable value for the parameter (including valid values, invalid values, and `NULL`).

2

The global commands are: [vkEnumerateInstanceVersion](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceVersion.html), [vkEnumerateInstanceExtensionProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceExtensionProperties.html), [vkEnumerateInstanceLayerProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkEnumerateInstanceLayerProperties.html), and [vkCreateInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateInstance.html). Dispatchable commands are all other commands which are not global.

3

The returned function pointer **must** only be called with a dispatchable object (the first parameter) that is `instance` or a child of `instance`, e.g. [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkQueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueue.html), or [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html).

4

An “available device extension” is a device extension supported by any physical device enumerated by `instance`.

5

Starting with Vulkan 1.2, `vkGetInstanceProcAddr` can resolve itself with a `NULL` instance pointer.

Valid Usage (Implicit)

- [](#VUID-vkGetInstanceProcAddr-instance-parameter)VUID-vkGetInstanceProcAddr-instance-parameter  
  If `instance` is not `NULL`, `instance` **must** be a valid [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html) handle
- [](#VUID-vkGetInstanceProcAddr-pName-parameter)VUID-vkGetInstanceProcAddr-pName-parameter  
  `pName` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[PFN\_vkVoidFunction](https://registry.khronos.org/vulkan/specs/latest/man/html/PFN_vkVoidFunction.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkInstance](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstance.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetInstanceProcAddr)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0