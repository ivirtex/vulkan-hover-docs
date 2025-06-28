# VkImportMemoryWin32HandleInfoKHR(3) Manual Page

## Name

VkImportMemoryWin32HandleInfoKHR - Import Win32 memory created on the same physical device



## [](#_c_specification)C Specification

To import memory from a Windows handle, add a [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryWin32HandleInfoKHR.html) structure to the `pNext` chain of the [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) structure.

The `VkImportMemoryWin32HandleInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_external_memory_win32
typedef struct VkImportMemoryWin32HandleInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkExternalMemoryHandleTypeFlagBits    handleType;
    HANDLE                                handle;
    LPCWSTR                               name;
} VkImportMemoryWin32HandleInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleType` is a [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value specifying the type of `handle` or `name`.
- `handle` is `NULL` or the external handle to import.
- `name` is `NULL` or a null-terminated UTF-16 string naming the payload to import.

## [](#_description)Description

Importing memory object payloads from Windows handles does not transfer ownership of the handle to the Vulkan implementation. For handle types defined as NT handles, the application **must** release handle ownership using the `CloseHandle` system call when the handle is no longer needed. For handle types defined as NT handles, the imported memory object holds a reference to its payload.

Note

Non-NT handle import operations do not add a reference to their associated payload. If the original object owning the payload is destroyed, all resources and handles sharing that payload will become invalid.

Applications **can** import the same payload into multiple instances of Vulkan, into the same instance from which it was exported, and multiple times into a given Vulkan instance. In all cases, each import operation **must** create a distinct `VkDeviceMemory` object.

Valid Usage

- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-09861)VUID-VkImportMemoryWin32HandleInfoKHR-handleType-09861  
  If `handleType` is not `0`, it **must** be supported for import, as reported by [VkExternalTensorPropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalTensorPropertiesARM.html), [VkExternalImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatProperties.html), or [VkExternalBufferProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalBufferProperties.html)
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handle-00659)VUID-VkImportMemoryWin32HandleInfoKHR-handle-00659  
  The memory from which `handle` was exported, or the memory named by `name` **must** have been created on the same underlying physical device as `device`
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00660)VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00660  
  If `handleType` is not `0`, it **must** be defined as an NT handle or a global share handle
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01439)VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01439  
  If `handleType` is not `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`, or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, `name` **must** be `NULL`
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01440)VUID-VkImportMemoryWin32HandleInfoKHR-handleType-01440  
  If `handleType` is not `0` and `handle` is `NULL`, `name` **must** name a valid memory resource of the type specified by `handleType`
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00661)VUID-VkImportMemoryWin32HandleInfoKHR-handleType-00661  
  If `handleType` is not `0` and `name` is `NULL`, `handle` **must** be a valid handle of the type specified by `handleType`
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handle-01441)VUID-VkImportMemoryWin32HandleInfoKHR-handle-01441  
  If `handle` is not `NULL`, `name` **must** be `NULL`
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handle-01518)VUID-VkImportMemoryWin32HandleInfoKHR-handle-01518  
  If `handle` is not `NULL`, it **must** obey any requirements listed for `handleType` in [external memory handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-memory-handle-types-compatibility)
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-name-01519)VUID-VkImportMemoryWin32HandleInfoKHR-name-01519  
  If `name` is not `NULL`, it **must** obey any requirements listed for `handleType` in [external memory handle types compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#external-memory-handle-types-compatibility)

Valid Usage (Implicit)

- [](#VUID-VkImportMemoryWin32HandleInfoKHR-sType-sType)VUID-VkImportMemoryWin32HandleInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR`
- [](#VUID-VkImportMemoryWin32HandleInfoKHR-handleType-parameter)VUID-VkImportMemoryWin32HandleInfoKHR-handleType-parameter  
  If `handleType` is not `0`, `handleType` **must** be a valid [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) value

## [](#_see_also)See Also

[VK\_KHR\_external\_memory\_win32](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory_win32.html), [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkImportMemoryWin32HandleInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0