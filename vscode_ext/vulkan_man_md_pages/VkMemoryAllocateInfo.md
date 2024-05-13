# VkMemoryAllocateInfo(3) Manual Page

## Name

VkMemoryAllocateInfo - Structure containing parameters of a memory
allocation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMemoryAllocateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkMemoryAllocateInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       allocationSize;
    uint32_t           memoryTypeIndex;
} VkMemoryAllocateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `allocationSize` is the size of the allocation in bytes.

- `memoryTypeIndex` is an index identifying a memory type from the
  `memoryTypes` array of the
  [VkPhysicalDeviceMemoryProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryProperties.html)
  structure.

## <a href="#_description" class="anchor"></a>Description

The internal data of an allocated device memory object **must** include
a reference to implementation-specific resources, referred to as the
memory object’s *payload*. Applications **can** also import and export
that internal data to and from device memory objects to share data
between Vulkan instances and other compatible APIs. A
`VkMemoryAllocateInfo` structure defines a memory import operation if
its `pNext` chain includes one of the following structures:

- [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html)
  with a non-zero `handleType` value

- [VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryFdInfoKHR.html) with a
  non-zero `handleType` value

- [VkImportMemoryHostPointerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryHostPointerInfoEXT.html)
  with a non-zero `handleType` value

- [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  with a non-`NULL` `buffer` value

- [VkImportMemoryZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryZirconHandleInfoFUCHSIA.html)
  with a non-zero `handleType` value

- [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)

- [VkImportScreenBufferInfoQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportScreenBufferInfoQNX.html) with a
  non-`NULL` `buffer` value

If the parameters define an import operation and the external handle
type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`,
`VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT`, or
`VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, `allocationSize` is
ignored. The implementation **must** query the size of these allocations
from the OS.

Whether device memory objects constructed via a memory import operation
hold a reference to their payload depends on the properties of the
handle type used to perform the import, as defined below for each valid
handle type. Importing memory **must** not modify the content of the
memory. Implementations **must** ensure that importing memory does not
enable the importing Vulkan instance to access any memory or resources
in other Vulkan instances other than that corresponding to the memory
object imported. Implementations **must** also ensure accessing imported
memory which has not been initialized does not allow the importing
Vulkan instance to obtain data from the exporting Vulkan instance or
vice-versa.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>How exported and imported memory is isolated is left to the
implementation, but applications should be aware that such isolation
<strong>may</strong> prevent implementations from placing multiple
exportable memory objects in the same physical or virtual page. Hence,
applications <strong>should</strong> avoid creating many small external
memory objects whenever possible.</p></td>
</tr>
</tbody>
</table>

Importing memory **must** not increase overall heap usage within a
system. However, it **must** affect the following per-process values:

- [VkPhysicalDeviceMaintenance3Properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMaintenance3Properties.html)::`maxMemoryAllocationCount`

- [VkPhysicalDeviceMemoryBudgetPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceMemoryBudgetPropertiesEXT.html)::`heapUsage`

When performing a memory import operation, it is the responsibility of
the application to ensure the external handles and their associated
payloads meet all valid usage requirements. However, implementations
**must** perform sufficient validation of external handles and payloads
to ensure that the operation results in a valid memory object which will
not cause program termination, device loss, queue stalls, or corruption
of other resources when used as allowed according to its allocation
parameters. If the external handle provided does not meet these
requirements, the implementation **must** fail the memory import
operation with the error code `VK_ERROR_INVALID_EXTERNAL_HANDLE`. If the
parameters define an export operation and the external handle type is
`VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
implementations **should** not strictly follow `memoryTypeIndex`.
Instead, they **should** modify the allocation internally to use the
required memory type for the application’s given usage. This is because
for an export operation, there is currently no way for the client to
know the memory type index before allocating.

Valid Usage

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-07897"
  id="VUID-VkMemoryAllocateInfo-allocationSize-07897"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-07897  
  If the parameters do not define an <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-import-operation"
  target="_blank" rel="noopener">import or export operation</a>,
  `allocationSize` **must** be greater than `0`

- <a href="#VUID-VkMemoryAllocateInfo-None-06657"
  id="VUID-VkMemoryAllocateInfo-None-06657"></a>
  VUID-VkMemoryAllocateInfo-None-06657  
  The parameters **must** not define more than one <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-import-operation"
  target="_blank" rel="noopener">import operation</a>

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-07899"
  id="VUID-VkMemoryAllocateInfo-allocationSize-07899"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-07899  
  If the parameters define an export operation and the handle type is
  not
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID` ,
  `allocationSize` **must** be greater than `0`

- <a href="#VUID-VkMemoryAllocateInfo-buffer-06380"
  id="VUID-VkMemoryAllocateInfo-buffer-06380"></a>
  VUID-VkMemoryAllocateInfo-buffer-06380  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html), and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`
  is present and non-NULL,
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)::`collection`
  and
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)::`index`
  must match
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)::`collection`
  and
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)::`index`,
  respectively, of the
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
  structure used to create the
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`

- <a href="#VUID-VkMemoryAllocateInfo-image-06381"
  id="VUID-VkMemoryAllocateInfo-image-06381"></a>
  VUID-VkMemoryAllocateInfo-image-06381  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html), and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  is present and non-NULL,
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)::`collection`
  and
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)::`index`
  must match
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)::`collection`
  and
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)::`index`,
  respectively, of the
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
  structure used to create the
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-06382"
  id="VUID-VkMemoryAllocateInfo-allocationSize-06382"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-06382  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
  `allocationSize` **must** match
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html)::`size` value
  retrieved by
  [vkGetImageMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements.html) or
  [vkGetBufferMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferMemoryRequirements.html)
  for image-based or buffer-based collections respectively

- <a href="#VUID-VkMemoryAllocateInfo-pNext-06383"
  id="VUID-VkMemoryAllocateInfo-pNext-06383"></a>
  VUID-VkMemoryAllocateInfo-pNext-06383  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html), the
  `pNext` chain **must** include a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with either its `image` or `buffer` field set to a value
  other than [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkMemoryAllocateInfo-image-06384"
  id="VUID-VkMemoryAllocateInfo-image-06384"></a>
  VUID-VkMemoryAllocateInfo-image-06384  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `image` **must** be
  created with a
  [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
  structure chained to its
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`pNext` pointer

- <a href="#VUID-VkMemoryAllocateInfo-buffer-06385"
  id="VUID-VkMemoryAllocateInfo-buffer-06385"></a>
  VUID-VkMemoryAllocateInfo-buffer-06385  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html) and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`buffer`
  is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the `buffer` **must** be
  created with a
  [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
  structure chained to its
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`pNext` pointer

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-06386"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-06386"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-06386  
  If the parameters define an import operation from an
  [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html),
  `memoryTypeIndex` **must** be from
  [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)
  as retrieved by
  [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html)

- <a href="#VUID-VkMemoryAllocateInfo-pNext-00639"
  id="VUID-VkMemoryAllocateInfo-pNext-00639"></a>
  VUID-VkMemoryAllocateInfo-pNext-00639  
  If the `pNext` chain includes a `VkExportMemoryAllocateInfo`
  structure, and any of the handle types specified in
  `VkExportMemoryAllocateInfo`::`handleTypes` require a dedicated
  allocation, as reported by
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
  in
  `VkExternalImageFormatProperties`::`externalMemoryProperties.externalMemoryFeatures`,
  or by
  [vkGetPhysicalDeviceExternalBufferProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceExternalBufferProperties.html)
  in
  `VkExternalBufferProperties`::`externalMemoryProperties.externalMemoryFeatures`,
  the `pNext` chain **must** include a `VkMemoryDedicatedAllocateInfo`
  or `VkDedicatedAllocationMemoryAllocateInfoNV` structure with either
  its `image` or `buffer` member set to a value other than
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkMemoryAllocateInfo-pNext-00640"
  id="VUID-VkMemoryAllocateInfo-pNext-00640"></a>
  VUID-VkMemoryAllocateInfo-pNext-00640  
  If the `pNext` chain includes a
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)
  structure, it **must** not include a
  [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html) or
  [VkExportMemoryWin32HandleInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoNV.html)
  structure

- <a href="#VUID-VkMemoryAllocateInfo-pNext-00641"
  id="VUID-VkMemoryAllocateInfo-pNext-00641"></a>
  VUID-VkMemoryAllocateInfo-pNext-00641  
  If the `pNext` chain includes a
  [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html)
  structure, it **must** not include a
  [VkImportMemoryWin32HandleInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoNV.html)
  structure

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-01742"
  id="VUID-VkMemoryAllocateInfo-allocationSize-01742"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-01742  
  If the parameters define an import operation, the external handle
  specified was created by the Vulkan API, and the external handle type
  is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`, then the values of
  `allocationSize` and `memoryTypeIndex` **must** match those specified
  when the payload being imported was created

- <a href="#VUID-VkMemoryAllocateInfo-None-00643"
  id="VUID-VkMemoryAllocateInfo-None-00643"></a>
  VUID-VkMemoryAllocateInfo-None-00643  
  If the parameters define an import operation and the external handle
  specified was created by the Vulkan API, the device mask specified by
  [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfo.html) **must**
  match the mask specified when the payload being imported was allocated

- <a href="#VUID-VkMemoryAllocateInfo-None-00644"
  id="VUID-VkMemoryAllocateInfo-None-00644"></a>
  VUID-VkMemoryAllocateInfo-None-00644  
  If the parameters define an import operation and the external handle
  specified was created by the Vulkan API, the list of physical devices
  that comprise the logical device passed to
  [vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html) **must** match the list of
  physical devices that comprise the logical device on which the payload
  was originally allocated

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-00645"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-00645"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-00645  
  If the parameters define an import operation and the external handle
  is an NT handle or a global share handle created outside of the Vulkan
  API, the value of `memoryTypeIndex` **must** be one of those returned
  by
  [vkGetMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandlePropertiesKHR.html)

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-01743"
  id="VUID-VkMemoryAllocateInfo-allocationSize-01743"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-01743  
  If the parameters define an import operation, the external handle was
  created by the Vulkan API, and the external handle type is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT` or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, then the values
  of `allocationSize` and `memoryTypeIndex` **must** match those
  specified when the payload being imported was created

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-00647"
  id="VUID-VkMemoryAllocateInfo-allocationSize-00647"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-00647  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`,
  `allocationSize` **must** match the size specified when creating the
  Direct3D 12 heap from which the payload was extracted

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-00648"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-00648"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-00648  
  If the parameters define an import operation and the external handle
  is a POSIX file descriptor created outside of the Vulkan API, the
  value of `memoryTypeIndex` **must** be one of those returned by
  [vkGetMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdPropertiesKHR.html)

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-01872"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-01872"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-01872  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-protectedMemory"
  target="_blank" rel="noopener"><code>protectedMemory</code></a>
  feature is not enabled, the `VkMemoryAllocateInfo`::`memoryTypeIndex`
  **must** not indicate a memory type that reports
  `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-01744"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-01744"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-01744  
  If the parameters define an import operation and the external handle
  is a host pointer, the value of `memoryTypeIndex` **must** be one of
  those returned by
  [vkGetMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryHostPointerPropertiesEXT.html)

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-01745"
  id="VUID-VkMemoryAllocateInfo-allocationSize-01745"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-01745  
  If the parameters define an import operation and the external handle
  is a host pointer, `allocationSize` **must** be an integer multiple of
  `VkPhysicalDeviceExternalMemoryHostPropertiesEXT`::`minImportedHostPointerAlignment`

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02805"
  id="VUID-VkMemoryAllocateInfo-pNext-02805"></a>
  VUID-VkMemoryAllocateInfo-pNext-02805  
  If the parameters define an import operation and the external handle
  is a host pointer, the `pNext` chain **must** not include a
  [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html)
  structure with either its `image` or `buffer` field set to a value
  other than [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02806"
  id="VUID-VkMemoryAllocateInfo-pNext-02806"></a>
  VUID-VkMemoryAllocateInfo-pNext-02806  
  If the parameters define an import operation and the external handle
  is a host pointer, the `pNext` chain **must** not include a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with either its `image` or `buffer` field set to a value
  other than [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-02383"
  id="VUID-VkMemoryAllocateInfo-allocationSize-02383"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-02383  
  If the parameters define an import operation and the external handle
  type is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  `allocationSize` **must** be the size returned by
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
  for the Android hardware buffer

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02384"
  id="VUID-VkMemoryAllocateInfo-pNext-02384"></a>
  VUID-VkMemoryAllocateInfo-pNext-02384  
  If the parameters define an import operation and the external handle
  type is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  and the `pNext` chain does not include a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure or
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the Android hardware buffer
  **must** have a `AHardwareBuffer_Desc`::`format` of
  `AHARDWAREBUFFER_FORMAT_BLOB` and a `AHardwareBuffer_Desc`::`usage`
  that includes `AHARDWAREBUFFER_USAGE_GPU_DATA_BUFFER`

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-02385"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-02385"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-02385  
  If the parameters define an import operation and the external handle
  type is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  `memoryTypeIndex` **must** be one of those returned by
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
  for the Android hardware buffer

- <a href="#VUID-VkMemoryAllocateInfo-pNext-01874"
  id="VUID-VkMemoryAllocateInfo-pNext-01874"></a>
  VUID-VkMemoryAllocateInfo-pNext-01874  
  If the parameters do not define an import operation, and the `pNext`
  chain includes a `VkExportMemoryAllocateInfo` structure with
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  included in its `handleTypes` member, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `image` not equal to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `allocationSize` **must**
  be `0`

- <a href="#VUID-VkMemoryAllocateInfo-pNext-07900"
  id="VUID-VkMemoryAllocateInfo-pNext-07900"></a>
  VUID-VkMemoryAllocateInfo-pNext-07900  
  If the parameters define an export operation, the handle type is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  and the `pNext` does not include a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure, `allocationSize` **must** be greater than `0`

- <a href="#VUID-VkMemoryAllocateInfo-pNext-07901"
  id="VUID-VkMemoryAllocateInfo-pNext-07901"></a>
  VUID-VkMemoryAllocateInfo-pNext-07901  
  If the parameters define an export operation, the handle type is
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`,
  and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `buffer` set to a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)
  object, `allocationSize` **must** be greater than `0`

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02386"
  id="VUID-VkMemoryAllocateInfo-pNext-02386"></a>
  VUID-VkMemoryAllocateInfo-pNext-02386  
  If the parameters define an import operation, the external handle is
  an Android hardware buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  with `image` that is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  Android hardware buffer’s
  [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/AHardwareBuffer.html)::`usage` **must** include at
  least one of `AHARDWAREBUFFER_USAGE_GPU_FRAMEBUFFER`,
  `AHARDWAREBUFFER_USAGE_GPU_SAMPLED_IMAGE` or
  `AHARDWAREBUFFER_USAGE_GPU_DATA_BUFFER`

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02387"
  id="VUID-VkMemoryAllocateInfo-pNext-02387"></a>
  VUID-VkMemoryAllocateInfo-pNext-02387  
  If the parameters define an import operation, the external handle is
  an Android hardware buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  with `image` that is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  format of `image` **must** be `VK_FORMAT_UNDEFINED` or the format
  returned by
  [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html)
  in
  [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html)::`format`
  for the Android hardware buffer

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02388"
  id="VUID-VkMemoryAllocateInfo-pNext-02388"></a>
  VUID-VkMemoryAllocateInfo-pNext-02388  
  If the parameters define an import operation, the external handle is
  an Android hardware buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `image` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the width, height, and array
  layer dimensions of `image` and the Android hardware buffer’s
  `AHardwareBuffer_Desc` **must** be identical

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02389"
  id="VUID-VkMemoryAllocateInfo-pNext-02389"></a>
  VUID-VkMemoryAllocateInfo-pNext-02389  
  If the parameters define an import operation, the external handle is
  an Android hardware buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `image` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the Android hardware
  buffer’s [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/AHardwareBuffer.html)::`usage` includes
  `AHARDWAREBUFFER_USAGE_GPU_MIPMAP_COMPLETE`, the `image` **must** have
  a complete mipmap chain

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02586"
  id="VUID-VkMemoryAllocateInfo-pNext-02586"></a>
  VUID-VkMemoryAllocateInfo-pNext-02586  
  If the parameters define an import operation, the external handle is
  an Android hardware buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `image` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and the Android hardware
  buffer’s [AHardwareBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/AHardwareBuffer.html)::`usage` does not
  include `AHARDWAREBUFFER_USAGE_GPU_MIPMAP_COMPLETE`, the `image`
  **must** have exactly one mipmap level

- <a href="#VUID-VkMemoryAllocateInfo-pNext-02390"
  id="VUID-VkMemoryAllocateInfo-pNext-02390"></a>
  VUID-VkMemoryAllocateInfo-pNext-02390  
  If the parameters define an import operation, the external handle is
  an Android hardware buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `image` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), each bit set in the usage of
  `image` **must** be listed in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-usage"
  target="_blank" rel="noopener">AHardwareBuffer Usage Equivalence</a>,
  and if there is a corresponding `AHARDWAREBUFFER_USAGE` bit listed
  that bit **must** be included in the Android hardware buffer’s
  `AHardwareBuffer_Desc`::`usage`

- <a href="#VUID-VkMemoryAllocateInfo-screenBufferImport-08941"
  id="VUID-VkMemoryAllocateInfo-screenBufferImport-08941"></a>
  VUID-VkMemoryAllocateInfo-screenBufferImport-08941  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`,
  [VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceExternalMemoryScreenBufferFeaturesQNX.html)::`screenBufferImport`
  **must** be enabled

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-08942"
  id="VUID-VkMemoryAllocateInfo-allocationSize-08942"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-08942  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`,
  `allocationSize` **must** be the size returned by
  [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)
  for the QNX Screen buffer

- <a href="#VUID-VkMemoryAllocateInfo-memoryTypeIndex-08943"
  id="VUID-VkMemoryAllocateInfo-memoryTypeIndex-08943"></a>
  VUID-VkMemoryAllocateInfo-memoryTypeIndex-08943  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_SCREEN_BUFFER_BIT_QNX`,
  `memoryTypeIndex` **must** be one of those returned by
  [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)
  for the QNX Screen buffer

- <a href="#VUID-VkMemoryAllocateInfo-pNext-08944"
  id="VUID-VkMemoryAllocateInfo-pNext-08944"></a>
  VUID-VkMemoryAllocateInfo-pNext-08944  
  If the parameters define an import operation, the external handle is a
  QNX Screen buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  with `image` that is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  QNX Screen’s buffer must be a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-screen-buffer-validity"
  target="_blank" rel="noopener">valid QNX Screen buffer</a>

- <a href="#VUID-VkMemoryAllocateInfo-pNext-08945"
  id="VUID-VkMemoryAllocateInfo-pNext-08945"></a>
  VUID-VkMemoryAllocateInfo-pNext-08945  
  If the parameters define an import operation, the external handle is
  an QNX Screen buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  with `image` that is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  format of `image` **must** be `VK_FORMAT_UNDEFINED` or the format
  returned by
  [vkGetScreenBufferPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetScreenBufferPropertiesQNX.html)
  in
  [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)::`format`
  for the QNX Screen buffer

- <a href="#VUID-VkMemoryAllocateInfo-pNext-08946"
  id="VUID-VkMemoryAllocateInfo-pNext-08946"></a>
  VUID-VkMemoryAllocateInfo-pNext-08946  
  If the parameters define an import operation, the external handle is a
  QNX Screen buffer, and the `pNext` chain includes a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure with `image` that is not
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the width, height, and array
  layer dimensions of `image` and the QNX Screen buffer’s
  `_screen_buffer` must be identical

- <a href="#VUID-VkMemoryAllocateInfo-opaqueCaptureAddress-03329"
  id="VUID-VkMemoryAllocateInfo-opaqueCaptureAddress-03329"></a>
  VUID-VkMemoryAllocateInfo-opaqueCaptureAddress-03329  
  If
  [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)::`opaqueCaptureAddress`
  is not zero, `VkMemoryAllocateFlagsInfo`::`flags` **must** include
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`

- <a href="#VUID-VkMemoryAllocateInfo-flags-03330"
  id="VUID-VkMemoryAllocateInfo-flags-03330"></a>
  VUID-VkMemoryAllocateInfo-flags-03330  
  If `VkMemoryAllocateFlagsInfo`::`flags` includes
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddressCaptureReplay"
  target="_blank"
  rel="noopener"><code>bufferDeviceAddressCaptureReplay</code></a>
  feature **must** be enabled

- <a href="#VUID-VkMemoryAllocateInfo-flags-03331"
  id="VUID-VkMemoryAllocateInfo-flags-03331"></a>
  VUID-VkMemoryAllocateInfo-flags-03331  
  If `VkMemoryAllocateFlagsInfo`::`flags` includes
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT`, the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-bufferDeviceAddress"
  target="_blank" rel="noopener"><code>bufferDeviceAddress</code></a>
  feature **must** be enabled

- <a href="#VUID-VkMemoryAllocateInfo-pNext-03332"
  id="VUID-VkMemoryAllocateInfo-pNext-03332"></a>
  VUID-VkMemoryAllocateInfo-pNext-03332  
  If the `pNext` chain includes a `VkImportMemoryHostPointerInfoEXT`
  structure,
  [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)::`opaqueCaptureAddress`
  **must** be zero

- <a href="#VUID-VkMemoryAllocateInfo-opaqueCaptureAddress-03333"
  id="VUID-VkMemoryAllocateInfo-opaqueCaptureAddress-03333"></a>
  VUID-VkMemoryAllocateInfo-opaqueCaptureAddress-03333  
  If the parameters define an import operation,
  [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)::`opaqueCaptureAddress`
  **must** be zero

- <a href="#VUID-VkMemoryAllocateInfo-None-04749"
  id="VUID-VkMemoryAllocateInfo-None-04749"></a>
  VUID-VkMemoryAllocateInfo-None-04749  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the
  value of `memoryTypeIndex` **must** be an index identifying a memory
  type from the `memoryTypeBits` field of the
  [VkMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryZirconHandlePropertiesFUCHSIA.html)
  structure populated by a call to
  [vkGetMemoryZirconHandlePropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryZirconHandlePropertiesFUCHSIA.html)

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-07902"
  id="VUID-VkMemoryAllocateInfo-allocationSize-07902"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-07902  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the
  value of `allocationSize` **must** be greater than `0`

- <a href="#VUID-VkMemoryAllocateInfo-allocationSize-07903"
  id="VUID-VkMemoryAllocateInfo-allocationSize-07903"></a>
  VUID-VkMemoryAllocateInfo-allocationSize-07903  
  If the parameters define an import operation and the external handle
  type is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the
  value of `allocationSize` **must** be less than or equal to the size
  of the VMO as determined by `zx_vmo_get_size`(`handle`) where `handle`
  is the VMO handle to the imported external memory

- <a href="#VUID-VkMemoryAllocateInfo-pNext-06780"
  id="VUID-VkMemoryAllocateInfo-pNext-06780"></a>
  VUID-VkMemoryAllocateInfo-pNext-06780  
  If the `pNext` chain includes a
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)
  structure, its `exportObjectType` member **must** be
  `VK_EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryAllocateInfo-sType-sType"
  id="VUID-VkMemoryAllocateInfo-sType-sType"></a>
  VUID-VkMemoryAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO`

- <a href="#VUID-VkMemoryAllocateInfo-pNext-pNext"
  id="VUID-VkMemoryAllocateInfo-pNext-pNext"></a>
  VUID-VkMemoryAllocateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html),
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html),
  [VkExportMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfoNV.html),
  [VkExportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoKHR.html),
  [VkExportMemoryWin32HandleInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoNV.html),
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html),
  [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html),
  [VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryFdInfoKHR.html),
  [VkImportMemoryHostPointerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryHostPointerInfoEXT.html),
  [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html),
  [VkImportMemoryWin32HandleInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoNV.html),
  [VkImportMemoryZirconHandleInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryZirconHandleInfoFUCHSIA.html),
  [VkImportMetalBufferInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMetalBufferInfoEXT.html),
  [VkImportScreenBufferInfoQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportScreenBufferInfoQNX.html),
  [VkMemoryAllocateFlagsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateFlagsInfo.html),
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html),
  [VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html),
  or
  [VkMemoryPriorityAllocateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPriorityAllocateInfoEXT.html)

- <a href="#VUID-VkMemoryAllocateInfo-sType-unique"
  id="VUID-VkMemoryAllocateInfo-sType-unique"></a>
  VUID-VkMemoryAllocateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique, with the exception of structures of type
  [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkAllocateMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAllocateMemory.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
