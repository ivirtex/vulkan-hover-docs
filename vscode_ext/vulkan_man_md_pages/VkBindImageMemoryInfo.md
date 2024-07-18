# VkBindImageMemoryInfo(3) Manual Page

## Name

VkBindImageMemoryInfo - Structure specifying how to bind an image to
memory



## <a href="#_c_specification" class="anchor"></a>C Specification

`VkBindImageMemoryInfo` contains members corresponding to the parameters
of [vkBindImageMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory.html).

The `VkBindImageMemoryInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkBindImageMemoryInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    VkDeviceMemory     memory;
    VkDeviceSize       memoryOffset;
} VkBindImageMemoryInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_bind_memory2
typedef VkBindImageMemoryInfo VkBindImageMemoryInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is the image to be attached to memory.

- `memory` is a [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object describing
  the device memory to attach.

- `memoryOffset` is the start offset of the region of `memory` which is
  to be bound to the image. The number of bytes returned in the
  `VkMemoryRequirements`::`size` member in `memory`, starting from
  `memoryOffset` bytes, will be bound to the specified image.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkBindImageMemoryInfo-image-07460"
  id="VUID-VkBindImageMemoryInfo-image-07460"></a>
  VUID-VkBindImageMemoryInfo-image-07460  
  `image` **must** not have been bound to a memory object

- <a href="#VUID-VkBindImageMemoryInfo-image-01045"
  id="VUID-VkBindImageMemoryInfo-image-01045"></a>
  VUID-VkBindImageMemoryInfo-image-01045  
  `image` **must** not have been created with any sparse memory binding
  flags

- <a href="#VUID-VkBindImageMemoryInfo-memoryOffset-01046"
  id="VUID-VkBindImageMemoryInfo-memoryOffset-01046"></a>
  VUID-VkBindImageMemoryInfo-memoryOffset-01046  
  `memoryOffset` **must** be less than the size of `memory`

- <a href="#VUID-VkBindImageMemoryInfo-image-01445"
  id="VUID-VkBindImageMemoryInfo-image-01445"></a>
  VUID-VkBindImageMemoryInfo-image-01445  
  If `image` requires a dedicated allocation (as reported by
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html) in
  [VkMemoryDedicatedRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedRequirements.html)::`requiresDedicatedAllocation`
  for `image`), `memory` **must** have been created with
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  equal to `image`

- <a href="#VUID-VkBindImageMemoryInfo-memory-02628"
  id="VUID-VkBindImageMemoryInfo-memory-02628"></a>
  VUID-VkBindImageMemoryInfo-memory-02628  
  If the
  [`dedicatedAllocationImageAliasing`](#features-dedicatedAllocationImageAliasing)
  feature is not enabled, and the `VkMemoryAllocateInfo` provided when
  `memory` was allocated included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure in its `pNext` chain, and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `image` **must**
  equal
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  and `memoryOffset` **must** be zero

- <a href="#VUID-VkBindImageMemoryInfo-memory-02629"
  id="VUID-VkBindImageMemoryInfo-memory-02629"></a>
  VUID-VkBindImageMemoryInfo-memory-02629  
  If the
  [`dedicatedAllocationImageAliasing`](#features-dedicatedAllocationImageAliasing)
  feature is enabled, and the `VkMemoryAllocateInfo` provided when
  `memory` was allocated included a
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)
  structure in its `pNext` chain, and
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  was not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then `memoryOffset`
  **must** be zero, and `image` **must** be either equal to
  [VkMemoryDedicatedAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryDedicatedAllocateInfo.html)::`image`
  or an image that was created using the same parameters in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html), with the exception that
  `extent` and `arrayLayers` **may** differ subject to the following
  restrictions: every dimension in the `extent` parameter of the image
  being bound **must** be equal to or smaller than the original image
  for which the allocation was created; and the `arrayLayers` parameter
  of the image being bound **must** be equal to or smaller than the
  original image for which the allocation was created

- <a href="#VUID-VkBindImageMemoryInfo-None-01901"
  id="VUID-VkBindImageMemoryInfo-None-01901"></a>
  VUID-VkBindImageMemoryInfo-None-01901  
  If image was created with the `VK_IMAGE_CREATE_PROTECTED_BIT` bit set,
  the image **must** be bound to a memory object allocated with a memory
  type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-VkBindImageMemoryInfo-None-01902"
  id="VUID-VkBindImageMemoryInfo-None-01902"></a>
  VUID-VkBindImageMemoryInfo-None-01902  
  If image was created with the `VK_IMAGE_CREATE_PROTECTED_BIT` bit not
  set, the image **must** not be bound to a memory object created with a
  memory type that reports `VK_MEMORY_PROPERTY_PROTECTED_BIT`

- <a href="#VUID-VkBindImageMemoryInfo-image-01050"
  id="VUID-VkBindImageMemoryInfo-image-01050"></a>
  VUID-VkBindImageMemoryInfo-image-01050  
  If `image` was created with
  [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** have been created with
  [VkDedicatedAllocationMemoryAllocateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationMemoryAllocateInfoNV.html)::`image`
  equal to an image handle created with identical creation parameters to
  `image` and `memoryOffset` **must** be zero

- <a href="#VUID-VkBindImageMemoryInfo-apiVersion-07921"
  id="VUID-VkBindImageMemoryInfo-apiVersion-07921"></a>
  VUID-VkBindImageMemoryInfo-apiVersion-07921  
  If the [VK_KHR_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dedicated_allocation.html)
  extension is not enabled,
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, and `image` was not created with
  [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`, `memory` **must** not have been allocated
  dedicated for a specific buffer or image

- <a href="#VUID-VkBindImageMemoryInfo-memory-02728"
  id="VUID-VkBindImageMemoryInfo-memory-02728"></a>
  VUID-VkBindImageMemoryInfo-memory-02728  
  If the value of
  [VkExportMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryAllocateInfo.html)::`handleTypes`
  used to allocate `memory` is not `0`, it **must** include at least one
  of the handles set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

- <a href="#VUID-VkBindImageMemoryInfo-memory-02989"
  id="VUID-VkBindImageMemoryInfo-memory-02989"></a>
  VUID-VkBindImageMemoryInfo-memory-02989  
  If `memory` was created by a memory import operation, that is not
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  with a non-`NULL` `buffer` value, the external handle type of the
  imported memory **must** also have been set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

- <a href="#VUID-VkBindImageMemoryInfo-memory-02990"
  id="VUID-VkBindImageMemoryInfo-memory-02990"></a>
  VUID-VkBindImageMemoryInfo-memory-02990  
  If `memory` was created with the
  [VkImportAndroidHardwareBufferInfoANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportAndroidHardwareBufferInfoANDROID.html)
  memory import operation with a non-`NULL` `buffer` value,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ANDROID_HARDWARE_BUFFER_BIT_ANDROID`
  **must** also have been set in
  [VkExternalMemoryImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryImageCreateInfo.html)::`handleTypes`
  when `image` was created

- <a
  href="#VUID-VkBindImageMemoryInfo-descriptorBufferCaptureReplay-08113"
  id="VUID-VkBindImageMemoryInfo-descriptorBufferCaptureReplay-08113"></a>
  VUID-VkBindImageMemoryInfo-descriptorBufferCaptureReplay-08113  
  If the `image` was created with the
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_BIT` bit set

- <a href="#VUID-VkBindImageMemoryInfo-image-09202"
  id="VUID-VkBindImageMemoryInfo-image-09202"></a>
  VUID-VkBindImageMemoryInfo-image-09202  
  If the `image` was created with the
  `VK_IMAGE_CREATE_DESCRIPTOR_BUFFER_CAPTURE_REPLAY_BIT_EXT` bit set,
  `memory` **must** have been allocated with the
  `VK_MEMORY_ALLOCATE_DEVICE_ADDRESS_CAPTURE_REPLAY_BIT` bit set

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01615"
  id="VUID-VkBindImageMemoryInfo-pNext-01615"></a>
  VUID-VkBindImageMemoryInfo-pNext-01615  
  If the `pNext` chain does not include a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, `memory` **must** have been allocated using one of the
  memory types allowed in the `memoryTypeBits` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)
  with `image`

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01616"
  id="VUID-VkBindImageMemoryInfo-pNext-01616"></a>
  VUID-VkBindImageMemoryInfo-pNext-01616  
  If the `pNext` chain does not include a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, `memoryOffset` **must** be an integer multiple of the
  `alignment` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)
  with `image`

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01617"
  id="VUID-VkBindImageMemoryInfo-pNext-01617"></a>
  VUID-VkBindImageMemoryInfo-pNext-01617  
  If the `pNext` chain does not include a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, the difference of the size of `memory` and `memoryOffset`
  **must** be greater than or equal to the `size` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)
  with the same `image`

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01618"
  id="VUID-VkBindImageMemoryInfo-pNext-01618"></a>
  VUID-VkBindImageMemoryInfo-pNext-01618  
  If the `pNext` chain includes a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, `image` **must** have been created with the
  `VK_IMAGE_CREATE_DISJOINT_BIT` bit set

- <a href="#VUID-VkBindImageMemoryInfo-image-07736"
  id="VUID-VkBindImageMemoryInfo-image-07736"></a>
  VUID-VkBindImageMemoryInfo-image-07736  
  If `image` was created with the `VK_IMAGE_CREATE_DISJOINT_BIT` bit
  set, then the `pNext` chain **must** include a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01619"
  id="VUID-VkBindImageMemoryInfo-pNext-01619"></a>
  VUID-VkBindImageMemoryInfo-pNext-01619  
  If the `pNext` chain includes a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, `memory` **must** have been allocated using one of the
  memory types allowed in the `memoryTypeBits` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)
  with `image` and where
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)::`planeAspect`
  corresponds to the
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)::`planeAspect`
  in the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure’s `pNext` chain

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01620"
  id="VUID-VkBindImageMemoryInfo-pNext-01620"></a>
  VUID-VkBindImageMemoryInfo-pNext-01620  
  If the `pNext` chain includes a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, `memoryOffset` **must** be an integer multiple of the
  `alignment` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)
  with `image` and where
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)::`planeAspect`
  corresponds to the
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)::`planeAspect`
  in the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure’s `pNext` chain

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01621"
  id="VUID-VkBindImageMemoryInfo-pNext-01621"></a>
  VUID-VkBindImageMemoryInfo-pNext-01621  
  If the `pNext` chain includes a
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)
  structure, the difference of the size of `memory` and `memoryOffset`
  **must** be greater than or equal to the `size` member of the
  [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements.html) structure returned
  from a call to
  [vkGetImageMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetImageMemoryRequirements2.html)
  with the same `image` and where
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html)::`planeAspect`
  corresponds to the
  [VkImagePlaneMemoryRequirementsInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImagePlaneMemoryRequirementsInfo.html)::`planeAspect`
  in the
  [VkImageMemoryRequirementsInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageMemoryRequirementsInfo2.html)
  structure’s `pNext` chain

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01626"
  id="VUID-VkBindImageMemoryInfo-pNext-01626"></a>
  VUID-VkBindImageMemoryInfo-pNext-01626  
  If the `pNext` chain includes a
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)
  structure, all instances of `memory` specified by
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)::`pDeviceIndices`
  **must** have been allocated

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01627"
  id="VUID-VkBindImageMemoryInfo-pNext-01627"></a>
  VUID-VkBindImageMemoryInfo-pNext-01627  
  If the `pNext` chain includes a
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)
  structure, and
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)::`splitInstanceBindRegionCount`
  is not zero, then `image` **must** have been created with the
  `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT` bit set

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01628"
  id="VUID-VkBindImageMemoryInfo-pNext-01628"></a>
  VUID-VkBindImageMemoryInfo-pNext-01628  
  If the `pNext` chain includes a
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)
  structure, all elements of
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)::`pSplitInstanceBindRegions`
  **must** be valid rectangles contained within the dimensions of
  `image`

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01629"
  id="VUID-VkBindImageMemoryInfo-pNext-01629"></a>
  VUID-VkBindImageMemoryInfo-pNext-01629  
  If the `pNext` chain includes a
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)
  structure, the union of the areas of all elements of
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html)::`pSplitInstanceBindRegions`
  that correspond to the same instance of `image` **must** cover the
  entire image

- <a href="#VUID-VkBindImageMemoryInfo-image-01630"
  id="VUID-VkBindImageMemoryInfo-image-01630"></a>
  VUID-VkBindImageMemoryInfo-image-01630  
  If `image` was created with a valid swapchain handle in
  [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html)::`swapchain`,
  then the `pNext` chain **must** include a
  [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html)
  structure containing the same swapchain handle

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01631"
  id="VUID-VkBindImageMemoryInfo-pNext-01631"></a>
  VUID-VkBindImageMemoryInfo-pNext-01631  
  If the `pNext` chain includes a
  [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html)
  structure, `memory` **must** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkBindImageMemoryInfo-pNext-01632"
  id="VUID-VkBindImageMemoryInfo-pNext-01632"></a>
  VUID-VkBindImageMemoryInfo-pNext-01632  
  If the `pNext` chain does not include a
  [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html)
  structure, `memory` **must** be a valid
  [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) handle

Valid Usage (Implicit)

- <a href="#VUID-VkBindImageMemoryInfo-sType-sType"
  id="VUID-VkBindImageMemoryInfo-sType-sType"></a>
  VUID-VkBindImageMemoryInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_INFO`

- <a href="#VUID-VkBindImageMemoryInfo-pNext-pNext"
  id="VUID-VkBindImageMemoryInfo-pNext-pNext"></a>
  VUID-VkBindImageMemoryInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html),
  [VkBindImageMemorySwapchainInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemorySwapchainInfoKHR.html),
  [VkBindImagePlaneMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImagePlaneMemoryInfo.html), or
  [VkBindMemoryStatusKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindMemoryStatusKHR.html)

- <a href="#VUID-VkBindImageMemoryInfo-sType-unique"
  id="VUID-VkBindImageMemoryInfo-sType-unique"></a>
  VUID-VkBindImageMemoryInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkBindImageMemoryInfo-image-parameter"
  id="VUID-VkBindImageMemoryInfo-image-parameter"></a>
  VUID-VkBindImageMemoryInfo-image-parameter  
  `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkBindImageMemoryInfo-commonparent"
  id="VUID-VkBindImageMemoryInfo-commonparent"></a>
  VUID-VkBindImageMemoryInfo-commonparent  
  Both of `image`, and `memory` that are valid handles of non-ignored
  parameters **must** have been created, allocated, or retrieved from
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBindImageMemory2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2.html),
[vkBindImageMemory2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBindImageMemory2KHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindImageMemoryInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
